package grpc

import (
	"crypto/tls"
	"crypto/x509"

	"github.com/richardcase/dockercoinsgo/certs"
	pb "github.com/richardcase/dockercoinsgo/hasher"
	"golang.org/x/net/context"
	ggrpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type GrpcHasherClient struct {
	ServerAddress  string
	CertFile       string
	KeyFile        string
	CAFile         string
	hasherKeyPair  *tls.Certificate
	hasherCertPool *x509.CertPool
}

func (c GrpcHasherClient) Hash(s string) (string, error) {
	conn, err := c.getConnection()
	if err != nil {
		return "", err
	}
	defer conn.Close()

	client := pb.NewHasherClient(conn)

	r, err := client.Hash(context.Background(), &pb.HashRequest{Message: s})
	if err != nil {
		return "", err
	}

	return r.HashedMessage, nil
}

func (c GrpcHasherClient) getConnection() (*ggrpc.ClientConn, error) {
	pair, err := certs.LoadCertificatesFromFile(c.CertFile, c.KeyFile)
	if err != nil {
		return nil, err
	}
	c.hasherKeyPair = &pair

	// Load the CA
	certPool, err := certs.CACertPoolFromFile(c.CAFile)
	if err != nil {
		return nil, err
	}
	c.hasherCertPool = certPool

	var opts []ggrpc.DialOption
	creds := credentials.NewClientTLSFromCert(c.hasherCertPool, c.ServerAddress)
	opts = append(opts, ggrpc.WithTransportCredentials(creds))
	conn, err := ggrpc.Dial(c.ServerAddress, opts...)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
