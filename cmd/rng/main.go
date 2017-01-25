package main

import (
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/richardcase/dockercoinsgo/certs"
	pb "github.com/richardcase/dockercoinsgo/rng"
)

var (
	rngKeyPair      *tls.Certificate
	rngCertPool     *x509.CertPool
	rngPort         *int
	rngCertHostname *string
	rngCertAddr     string
	rngLocalAddr    string
	rngPortAddr     string
	rngCert         *string
	rngKey          *string
	rngCA           *string
	shutdownDelay   *int
)

type server struct {
	isShuttingDown bool
}

func (s *server) GenerateRandom(ctx context.Context, in *pb.RngRequest) (*pb.RngResponse, error) {
	generated, err := generateRandomString(int(in.Length))
	if err != nil {
		log.Fatalf("Error generating random string: %v", err)
	}

	return &pb.RngResponse{Random: generated}, nil
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
	rngPort := flag.Int("port", 50052, "The port number to expose the server on")
	rngCertHostname := flag.String("certhostname", "localhost", "The published hostname of the service used for the certs")
	rngCert := flag.String("cert", "", "[Required]. Path to the certificate for the file")
	rngKey := flag.String("key", "", "[Required]. Path to the certificate key")
	rngCA := flag.String("ca", "", "[Required]. Path to the CA file")
	shutdownDelay = flag.Int("shut-delay", 4, "The delay in seconds to wait when sutting down")

	flag.Parse()

	if *rngCert == "" {
		fmt.Printf("You must supply a value for the 'cert' flag\n")
		os.Exit(1)
	}
	if *rngKey == "" {
		fmt.Printf("You must supply a value for the 'key' flag\n")
		os.Exit(1)
	}
	if *rngCA == "" {
		fmt.Printf("You must supply a value for the 'ca' flag\n")
		os.Exit(1)
	}
	/*if _, err := os.Stat(*rngCert); err == nil {
		fmt.Printf("Certificate file doesn't exist: %s\n", *rngCert)
		os.Exit(2)
	}
	if _, err := os.Stat(*rngKey); err == nil {
		fmt.Printf("Key file doesn't exist: %s\n", *rngKey)
		os.Exit(2)
	}*/

	rngCertAddr = fmt.Sprintf("%s:%d", *rngCertHostname, *rngPort)
	rngLocalAddr = fmt.Sprintf("localhost:%d", *rngPort)
	rngPortAddr = fmt.Sprintf(":%d", *rngPort)

	pair, err := certs.LoadCertificatesFromFile(*rngCert, *rngKey)
	if err != nil {
		fmt.Printf("Failed loading X509 key pair: %v\n", err)
		os.Exit(3)
	}
	rngKeyPair = &pair

	//Load the CA file
	certPool, err := certs.CACertPoolFromFile(*rngCA)
	if err != nil {
		fmt.Printf("Failed to load CA file: %s\n", rngCA)
		os.Exit(4)
	}
	rngCertPool = certPool

	serv := &server{isShuttingDown: false}

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Printf("SIGTERM received: Signalling shutdown\n")
		serv.SignalShutdown()
	}()

	opts := []grpc.ServerOption{
		grpc.Creds(credentials.NewClientTLSFromCert(rngCertPool, rngCertAddr))}

	s := grpc.NewServer(opts...)
	pb.RegisterRngServer(s, serv)
	ctx := context.Background()

	dcreds := credentials.NewTLS(&tls.Config{
		ServerName: rngCertAddr,
		RootCAs:    rngCertPool,
	})
	dopts := []grpc.DialOption{grpc.WithTransportCredentials(dcreds)}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", serv.HealthCheck)

	gwmux := runtime.NewServeMux()
	err = pb.RegisterRngHandlerFromEndpoint(ctx, gwmux, rngLocalAddr, dopts)
	if err != nil {
		fmt.Printf("serve: %v\n", err)
		return
	}

	mux.Handle("/", gwmux)

	lis, err := net.Listen("tcp", rngPortAddr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
		return
	}

	srv := &http.Server{
		Addr:    rngPortAddr,
		Handler: grpcHandlerFunc(s, mux),
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{*rngKeyPair},
			NextProtos:   []string{"h2"},
		},
	}

	fmt.Printf("RNG grpc on port: %d\n", *rngPort)

	if err := srv.Serve(tls.NewListener(lis, srv.TLSConfig)); err != nil {
		log.Fatalf("Failed t serve: %v", err)
	}

}

func generateRandomString(length int) (string, error) {
	b, err := generateRandomBytes(length)
	return base64.URLEncoding.EncodeToString(b), err

}

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
