package main

import (
	"fmt"
	"hash/fnv"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"net/http"
	"strings"

	"crypto/tls"
	"crypto/x509"

	"flag"

	"time"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/richardcase/dockercoinsgo/certs"
	pb "github.com/richardcase/dockercoinsgo/hasher"
)

var (
	hasherKeyPair  *tls.Certificate
	hasherCertPool *x509.CertPool
	hasherPort     *int
	hasherHostname *string
	hasherAddr     string
	hasherCert     *string
	hasherKey      *string
	hasherCA       *string
	shutdownDelay  *int
)

type server struct {
	isShuttingDown bool
}

func (s *server) Hash(ctx context.Context, in *pb.HashRequest) (*pb.HashResponse, error) {
	h := fnv.New32a()
	h.Write([]byte(in.Message))
	hashOutput := fmt.Sprintf("%d", h.Sum32())

	return &pb.HashResponse{HashedMessage: hashOutput}, nil
}

func (s *server) SignalShutdown() {
	s.isShuttingDown = true
	delay := time.Duration(*shutdownDelay) * time.Second
	timer := time.NewTimer(delay)
	go func() {
		<-timer.C
		fmt.Printf("Shutdown delay expired. Shutting down.\n")
		os.Exit(0)
	}()
}

func (s *server) HealthCheck(w http.ResponseWriter, r *http.Request) {
	if s.isShuttingDown == false {
		//NOTE: you would do other checks here if we are not shutting down
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Health check: OK"))
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Health check: NOT OK"))
	}
}

func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	})
}

func main() {
	hasherPort := flag.Int("port", 50051, "The port number to expose the server on")
	hasherHostname := flag.String("hostname", "localhost", "The published hostname of the service")
	hasherCert := flag.String("cert", "", "[Required]. Path to the certificate for the file")
	hasherKey := flag.String("key", "", "[Required]. Path to the certificate key")
	hasherCA := flag.String("ca", "", "[Required]. Path to the CA file")
	shutdownDelay = flag.Int("shut-delay", 4, "The delay in seconds to wait when sutting down")

	flag.Parse()

	if *hasherCert == "" {
		fmt.Printf("You must supply a value for the 'cert' flag\n")
		os.Exit(1)
	}
	if *hasherKey == "" {
		fmt.Printf("You must supply a value for the 'key' flag\n")
		os.Exit(1)
	}
	if *hasherCA == "" {
		fmt.Printf("You must supply a value for the 'ca' flag\n")
		os.Exit(1)
	}
	/*if _, err := os.Stat(*hasherCert); err == nil {
		fmt.Printf("Certificate file doesn't exist: %s\n", *hasherCert)
		os.Exit(2)
	}
	if _, err := os.Stat(*hasherKey); err == nil {
		fmt.Printf("Key file doesn't exist: %s\n", *hasherKey)
		os.Exit(2)
	}*/

	hasherAddr = fmt.Sprintf("%s:%d", *hasherHostname, *hasherPort)

	pair, err := certs.LoadCertificatesFromFile(*hasherCert, *hasherKey)
	if err != nil {
		fmt.Printf("Failed loading X509 key pair: %v\n", err)
		os.Exit(3)
	}
	hasherKeyPair = &pair

	//Load the CA file
	certPool, err := certs.CACertPoolFromFile(*hasherCA)
	if err != nil {
		fmt.Printf("Failed to load CA file: %s\n", hasherCA)
		os.Exit(4)
	}
	hasherCertPool = certPool

	serv := &server{isShuttingDown: false}

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Printf("SIGTERM received: Signalling shutdown\n")
		serv.SignalShutdown()
	}()

	opts := []grpc.ServerOption{
		grpc.Creds(credentials.NewClientTLSFromCert(hasherCertPool, hasherAddr))}

	s := grpc.NewServer(opts...)
	pb.RegisterHasherServer(s, serv)
	ctx := context.Background()

	dcreds := credentials.NewTLS(&tls.Config{
		ServerName: hasherAddr,
		RootCAs:    hasherCertPool,
	})
	dopts := []grpc.DialOption{grpc.WithTransportCredentials(dcreds)}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", serv.HealthCheck)

	gwmux := runtime.NewServeMux()
	err = pb.RegisterHasherHandlerFromEndpoint(ctx, gwmux, hasherAddr, dopts)
	if err != nil {
		fmt.Printf("serve: %v\n", err)
		return
	}

	mux.Handle("/", gwmux)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *hasherPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
		return
	}

	srv := &http.Server{
		Addr:    hasherAddr,
		Handler: grpcHandlerFunc(s, mux),
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{*hasherKeyPair},
			NextProtos:   []string{"h2"},
		},
	}

	fmt.Printf("Hasher grpc on port: %d\n", *hasherPort)

	if err := srv.Serve(tls.NewListener(lis, srv.TLSConfig)); err != nil {
		log.Fatalf("Failed t serve: %v", err)
	}

}
