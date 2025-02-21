package handlers

import (
	"context"
	validate "hostsetup-service/internal/service/entity"
	hsv1 "hostsetup-service/protos/gen/hostsetup"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
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
	ListDNSServers(ctx context.Context, dnsServers *[]string) error
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

	if err := validate.HostnameValidate(hostname); err != nil {
		return &hsv1.SuccessResponse{Success: false}, err
	}

	if err := s.changeHostname.SetHostname(ctx, hostname); err != nil {
		return &hsv1.SuccessResponse{Success: false}, err
	}

	return &hsv1.SuccessResponse{Success: true}, nil
}

func (s *serverAPI) ListDNSServers(
	ctx context.Context,
	req *emptypb.Empty,
) (*hsv1.ListDNSServersResponse, error) {
	dnsServers := make([]string, 0, 1)
	if err := s.setupDNS.ListDNSServers(ctx, &dnsServers); err != nil {
		return &hsv1.ListDNSServersResponse{}, err
	}

	return &hsv1.ListDNSServersResponse{List: dnsServers}, nil
}

func (s *serverAPI) AddDNSServer(
	ctx context.Context,
	req *hsv1.DNSServerRequest,
) (*hsv1.SuccessResponse, error) {
	dnsServer := req.DnsServer

	if err := validate.IPValidate(dnsServer); err != nil {
		return &hsv1.SuccessResponse{Success: false}, err
	}

	if err := s.setupDNS.AddDNSServer(ctx, dnsServer); err != nil {
		return &hsv1.SuccessResponse{Success: false}, err
	}

	return &hsv1.SuccessResponse{Success: true}, nil
}

func (s *serverAPI) DeleteDNSServer(
	ctx context.Context,
	req *hsv1.DNSServerRequest,
) (*hsv1.SuccessResponse, error) {
	dnsServer := req.DnsServer

	if err := validate.IPValidate(dnsServer); err != nil {
		return &hsv1.SuccessResponse{Success: false}, err
	}

	if err := s.setupDNS.DeleteDNSServer(ctx, dnsServer); err != nil {
		return &hsv1.SuccessResponse{Success: false}, err
	}

	return &hsv1.SuccessResponse{Success: true}, nil
}
