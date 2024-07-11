package app

import (
	hsv1 "hostsetup-service/protos/gen/hostsetup"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewClient(target string) (hsv1.HostSetupClient, error) {
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := hsv1.NewHostSetupClient(conn)

	return client, nil
}
