package setupping

import "context"

type Setupping struct{}

func New() *Setupping {
	return &Setupping{}
}

func (sp *Setupping) SetHostname(ctx context.Context, hostname string) (err error) {
	panic("unimplemented")
}

func (sp *Setupping) ListDNSServers(ctx context.Context) (err error) {
	panic("unimplemented")
}

func (sp *Setupping) AddDNSServer(ctx context.Context, dnsServer string) (err error) {
	panic("unimplemented")
}

func (sp *Setupping) DeleteDNSServer(ctx context.Context, dnsServer string) (err error) {
	panic("unimplemented")
}
