package grpc

import (
	"crypto/tls"
	"crypto/x509"

	"github.com/richardcase/dockercoinsgo/certs"
	pb "github.com/richardcase/dockercoinsgo/rng"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type GrpcRngClient struct {
	ServerAddress string
	CertFile      string
	KeyFile       string
	CAFile        string
	rngKeyPair    *tls.Certificate
	rngCertPool   *x509.CertPool
}

func (c GrpcRngClient) GenerateRandomString(length int32) (string, error) {
	conn, err := c.getConnection()
	if err != nil {
		return "", err
	}
	defer conn.Close()

	client := pb.NewRngClient(conn)

	r, err := client.GenerateRandom(context.Background(), &pb.RngRequest{Length: length})
	if err != nil {
		return "", err
	}

	return r.Random, nil
}

func (c GrpcRngClient) getConnection() (*grpc.ClientConn, error) {
	pair, err := certs.LoadCertificatesFromFile(c.CertFile, c.KeyFile)
	if err != nil {
		return nil, err
	}
	c.rngKeyPair = &pair

	// Load the CA
	certPool, err := certs.CACertPoolFromFile(c.CAFile)
	if err != nil {
		return nil, err
	}
	c.rngCertPool = certPool

	var opts []grpc.DialOption
	creds := credentials.NewClientTLSFromCert(c.rngCertPool, c.ServerAddress)
	opts = append(opts, grpc.WithTransportCredentials(creds))

	conn, err := grpc.Dial(c.ServerAddress, opts...)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
