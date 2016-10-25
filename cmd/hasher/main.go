package main

import (
	"crypto/sha256"
	"fmt"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/richardcase/dockercoinsgo/hasher"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) Hash(ctx context.Context, in *pb.HashRequest) (*pb.HashResponse, error) {
	hash := sha256.New()
	hash.Write([]byte(in.Message))
	hashOutput := fmt.Sprintf("%s", sha256.Sum256(nil))

	fmt.Println(hashOutput)

	return &pb.HashResponse{HashedMessage: hashOutput}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHasherServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed t serve: %v", err)
	}

}
