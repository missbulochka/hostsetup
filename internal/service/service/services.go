package services

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const resolvConfPrefix = "nameserver"

func ReadDNSServers(file *os.File, dnsServers *[]string) error {
	const op = "hostsetup: services.ReadDNSServers"

	log.Printf("reading dns servers")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), resolvConfPrefix) {
			*dnsServers = append(*dnsServers, scanner.Text())
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("%s:%w", op, err)
	}

	log.Printf("dns servers successfully read")
	return nil
}
