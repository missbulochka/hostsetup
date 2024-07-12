package main

import (
	"fmt"
	"hostsetup-service/pkg/config"
	hsv1 "hostsetup-service/protos/gen/hostsetup"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var gwAddr = fmt.Sprintf("%s:%s", "0.0.0.0", "8083")

func run() error {
	cfg := config.MustLoadConfig()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	hsAddr := fmt.Sprintf("%s:%s", cfg.GRPCServer, cfg.GRPCPort)

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := hsv1.RegisterHostSetupHandlerFromEndpoint(ctx, mux, hsAddr, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(gwAddr, mux)
}

func main() {
	if err := run(); err != nil {
		log.Println(err)
	}
}
