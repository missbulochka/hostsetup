package app

import (
	"context"
	hsv1 "hostsetup-service/protos/gen/hostsetup"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var Socket string

func NewClient(ctx context.Context) (*hsv1.HostSetupClient, error) {
	conn, err := grpc.Dial(Socket, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := hsv1.NewHostSetupClient(conn)

	return &client, nil
}
