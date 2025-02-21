package grpcapp

import (
	"fmt"
	setupgrpc "hostsetup-service/internal/service/handler/grpc"
	"log"
	"net"

	"google.golang.org/grpc"
)

type App struct {
	gRPCServer *grpc.Server
	server     string
	port       string
}

func New(server, port string,
	changeHostname setupgrpc.ChangeHostname,
	setupDNS setupgrpc.SetupDNS,
) *App {
	gRPCServer := grpc.NewServer()

	setupgrpc.Register(gRPCServer, changeHostname, setupDNS)

	return &App{
		gRPCServer: gRPCServer,
		server:     server,
		port:       port,
	}
}

func (a *App) MustStart() {
	if err := a.startServer(); err != nil {
		panic(err)
	}
}

func (a *App) startServer() error {
	const op = "hostsetup: grpcapp.startServer"

	l, err := net.Listen("tcp", fmt.Sprintf("%s:%s", a.server, a.port))
	if err != nil {
		return fmt.Errorf("%s:%w", op, err)
	}

	log.Printf("grpc server is running and listening at %s:%s", a.server, a.port)

	if err = a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s:%w", op, err)
	}

	return nil
}

func (a *App) Stop() {
	log.Printf("stopping grpc server %s:%s", a.server, a.port)

	a.gRPCServer.GracefulStop()
}
