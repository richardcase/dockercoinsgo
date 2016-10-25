package grpc

import (
	pb "github.com/richardcase/dockercoinsgo/hasher"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type GrpcHasherClient struct {
	ServerAddress string
}

func (c GrpcHasherClient) Hash(s string) (string, error) {
	conn, err := grpc.Dial(c.ServerAddress, grpc.WithInsecure())
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
