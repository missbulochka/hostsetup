package main

import (
	"hostsetup-service/internal/service/app"
	"hostsetup-service/pkg/config"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoadConfig()

	hostsetupService := app.New(cfg.GRPCServer, cfg.GRPCPort)

	go func() {
		hostsetupService.GRPCSrv.MustStart()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop
	log.Printf("shutting down hostsetup service by %v", sign)
	hostsetupService.GRPCSrv.Stop()
	log.Printf("hostsetup service stopped")
}
