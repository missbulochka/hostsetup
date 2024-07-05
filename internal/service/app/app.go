package app

import (
	grpcapp "hostsetup-service/internal/service/app/grpc"
	spuc "hostsetup-service/internal/service/usecase"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(server, port string) *App {
	setupping := spuc.New()
	grpcApp := grpcapp.New(server, port, setupping, setupping)

	return &App{
		GRPCSrv: grpcApp,
	}
}
