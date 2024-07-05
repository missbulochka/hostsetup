package handlers

import (
	"context"
	validator "hostsetup-service/internal/service/entity"
	hsv1 "hostsetup-service/protos/gen/hostsetup"

	"google.golang.org/grpc"
)

type serverAPI struct {
	hsv1.UnimplementedHostSetupServer
	changeHostname ChangeHostname
	setupDNS       SetupDNS
}

type ChangeHostname interface {
	SetHostname(ctx context.Context, hostname string) (err error)
}

type SetupDNS interface {
	ListDNSServers(ctx context.Context) error
	AddDNSServer(ctx context.Context, dnsServer string) error
	DeleteDNSServer(ctx context.Context, dnsServer string) error
}

func Register(
	server *grpc.Server,
	changeHostname ChangeHostname,
	setupDNS SetupDNS,
) {
	hsv1.RegisterHostSetupServer(
		server,
		&serverAPI{
			changeHostname: changeHostname,
			setupDNS:       setupDNS,
		},
	)
}

func (s *serverAPI) SetHostname(
	ctx context.Context,
	req *hsv1.HostnameRequest,
) (*hsv1.SuccessResponse, error) {
	hostname := req.Name

	if err := validator.HostnameValidate(hostname); err != nil {
		return &hsv1.SuccessResponse{Success: false}, err
	}

	if err := s.changeHostname.SetHostname(ctx, hostname); err != nil {
		return &hsv1.SuccessResponse{Success: false}, err
	}

	return &hsv1.SuccessResponse{Success: true}, nil
}

func (s *serverAPI) ListDNSServers(
	ctx context.Context,
	req *hsv1.EmptyRequest,
) (*hsv1.ListDNSServersResponse, error) {
	panic("implement me!")
}

func (s *serverAPI) AddDNSServer(
	ctx context.Context,
	req *hsv1.DNSServerRequest,
) (*hsv1.SuccessResponse, error) {
	panic("implement me!")
}

func (s *serverAPI) DeleteDNSServer(
	ctx context.Context,
	req *hsv1.DNSServerRequest,
) (*hsv1.SuccessResponse, error) {
	panic("implement me")
}
