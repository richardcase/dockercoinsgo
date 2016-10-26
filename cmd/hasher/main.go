package main

import (
	"fmt"
	"hash/fnv"
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
	h := fnv.New32a()
	h.Write([]byte(in.Message))
	hashOutput := fmt.Sprintf("%d", h.Sum32())

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
