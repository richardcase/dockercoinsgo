package main

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/richardcase/dockercoinsgo/rng"
)

const (
	port = ":50052"
)

type server struct{}

func (s *server) GenerateRandom(ctx context.Context, in *pb.RngRequest) (*pb.RngResponse, error) {
	generated, err := generateRandomString(int(in.Length))
	if err != nil {
		log.Fatalf("Error generating random string: %v", err)
	}

	return &pb.RngResponse{Random: generated}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterRngServer(s, &server{})
	if err := s.Serve(lis); err != nil {
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
