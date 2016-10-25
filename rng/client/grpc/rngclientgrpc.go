package grpc

import (
	pb "github.com/richardcase/dockercoinsgo/rng"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type GrpcRngClient struct {
	ServerAddress string
}

func (c GrpcRngClient) GenerateRandomString(length int32) (string, error) {
	conn, err := grpc.Dial(c.ServerAddress, grpc.WithInsecure())
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
