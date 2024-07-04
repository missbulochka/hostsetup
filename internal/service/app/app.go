package app

import (
	grpcapp "hostsetup-service/internal/service/app/grpc"
	"hostsetup-service/internal/service/config"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(cfg *config.GRPCConfig) *App {
	grpcApp := grpcapp.New(cfg.GRPCServer, cfg.GRPCPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}
