package services

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func VerifyHostname(hostname string) error {
	const op = "hostsetup: services.VerifyHostname "

	currentHostname, err := os.Hostname()
	if err != nil {
		log.Printf("%s: %v", op, err)
		return fmt.Errorf("%s:%w", op, err)
	}
	if currentHostname != hostname {
		log.Printf("%s: %v", op, err)
		return fmt.Errorf("%s:%s", op, "failed to change hostname")
	}

	return nil
}

func DNSServerExists(file *os.File, dnsServerStr string) (bool, error) {
	const op = "hostsetup: services.DNSServerExists"

	dnsServerStr = strings.Trim(dnsServerStr, " \n")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Trim(line, " \n")
		if line == dnsServerStr {
			return true, nil
		}
	}
	if err := scanner.Err(); err != nil {
		log.Printf("%s: %v", op, err)
		return false, fmt.Errorf("%s:%w", op, err)
	}

	return false, nil
}
