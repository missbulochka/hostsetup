package setupping

import (
	"bufio"
	"context"
	"fmt"
	services "hostsetup-service/internal/service/service"
	"log"
	"os"
	"os/exec"
	"strings"
)

const (
	pwdToResolvConf  string = "/etc/resolv.conf"
	resolvConfPrefix string = "nameserver "
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
		log.Printf("%s: %v", op, err)
		return fmt.Errorf("%s:%w", op, err)
	}

	if err := services.VerifyHostname(hostname); err != nil {
		return err
	}
	log.Printf("hostname successfully set")

	return nil
}

// ListDNSServers return dns servers
func (sp *Setupping) ListDNSServers(ctx context.Context, dnsServers *[]string) error {
	const op = "hostsetup: setupping.ListDNSServers"

	file, err := os.Open(pwdToResolvConf)
	if err != nil {
		log.Printf("%s: %v", op, err)
		return fmt.Errorf("%s:%w", op, err)
	}
	defer file.Close()

	log.Printf("reading dns servers")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " ")
		if strings.HasPrefix(line, resolvConfPrefix) {
			*dnsServers = append(*dnsServers, strings.TrimPrefix(line, resolvConfPrefix))
		}
	}
	if err := scanner.Err(); err != nil {
		log.Printf("%s: %v", op, err)
		return fmt.Errorf("%s:%w", op, err)
	}
	log.Printf("dns servers successfully read")

	return nil
}

// AddDNSServer add new dns server in the system
func (sp *Setupping) AddDNSServer(ctx context.Context, dnsServer string) error {
	const op = "hostsetup: setupping.AddDNSServer"

	file, err := os.OpenFile(pwdToResolvConf, os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Printf("%s: %v", op, err)
		return fmt.Errorf("%s:%w", op, err)
	}
	defer file.Close()

	resolverStringToAdd := resolvConfPrefix + dnsServer + "\n"
	exist, err := services.DNSServerExists(file, resolverStringToAdd)
	if err != nil {
		log.Printf("%s: %v", op, err)
		return fmt.Errorf("%s:%s", op, err)
	}
	if exist {
		return fmt.Errorf("%s:%s", op, "dns server already exist")
	}

	log.Printf("adding dns server")
	if _, err := file.WriteString(resolverStringToAdd); err != nil {
		log.Printf("%s: %v", op, err)
		return fmt.Errorf("%s:%w", op, err)
	}
	log.Printf("dns server successfully added")

	return nil
}

// DeleteDNSServer delete dns server from the system
func (sp *Setupping) DeleteDNSServer(ctx context.Context, dnsServer string) error {
	const op = "hostsetup: setupping.DeleteDNSServer"

	file, err := os.OpenFile(pwdToResolvConf, os.O_RDWR, 0666)
	if err != nil {
		log.Printf("%s: %v", op, err)
		return fmt.Errorf("%s:%w", op, err)
	}
	defer file.Close()

	log.Printf("delining dns server")
	// TODO: копирование текущих настроек в резервный файл
	resolverStringToDelete := resolvConfPrefix + dnsServer
	var newFile string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if resolverStringToDelete != strings.Trim(line, " ") {
			newFile = newFile + line + "\n"
		}
	}
	if err := scanner.Err(); err != nil {
		log.Printf("%s: %v", op, err)
		return fmt.Errorf("%s:%w", op, err)
	}

	if err := os.WriteFile(pwdToResolvConf, []byte(newFile), 0666); err != nil {
		log.Printf("%s: %v", op, err)
		return fmt.Errorf("%s:%w", op, err)
	}
	log.Printf("dns server successfully deleted")

	return nil
}
