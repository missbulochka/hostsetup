package setupping

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

const (
	pwdToResolvConf  = "/etc/resolv.conf"
	resolvConfPrefix = "nameserver "
)

type Setupping struct{}

func New() *Setupping {
	return &Setupping{}
}

// SetHostname set new hostname in the system
func (sp *Setupping) SetHostname(ctx context.Context, hostname string) error {
	const op = "hostsetup: setupping.SetHostname"

	log.Printf("setting hostname")
	cmd := exec.Command("hostname", hostname)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%s:%w", op, err)
	}

	currentHostname, err := os.Hostname()
	if err != nil {
		return fmt.Errorf("%s:%w", op, err)
	}
	if currentHostname != hostname {
		return fmt.Errorf("%s:%s", op, "failed to change hostname")
	}
	log.Printf("hostname set")

	return nil
}

// ListDNSServers return dns servers
func (sp *Setupping) ListDNSServers(ctx context.Context, dnsServers *[]string) error {
	const op = "hostsetup: setupping.ListDNSServers"

	file, err := os.Open(pwdToResolvConf)
	if err != nil {
		return fmt.Errorf("%s:%w", op, err)
	}
	defer file.Close()

	log.Printf("reading dns servers")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, resolvConfPrefix) {
			*dnsServers = append(*dnsServers, strings.TrimPrefix(line, resolvConfPrefix))
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("%s:%w", op, err)
	}
	log.Printf("dns servers successfully read")

	return nil
}

func (sp *Setupping) AddDNSServer(ctx context.Context, dnsServer string) error {
	panic("unimplemented")
}

func (sp *Setupping) DeleteDNSServer(ctx context.Context, dnsServer string) error {
	panic("unimplemented")
}
