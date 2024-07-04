package handlers

import (
	"context"
	hsv1 "hostsetup-service/protos/gen/hostsetup"
)

type Hostsetup struct {
	hsv1.UnimplementedHostSetupServer
	hsUseCase HostsetupUseCase
}

type HostsetupUseCase interface {
	SetHostname(ctx context.Context, hostname string) (err error)
	ListDNSServers(ctx context.Context)
}
