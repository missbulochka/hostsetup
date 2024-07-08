package setupping

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

func DNSServerExists(file *os.File, resolverString string) (bool, error) {
	const op = "hostsetup: services.DNSServerExists"

	resolverString = strings.Trim(resolverString, " ")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " ")
		if line == resolverString {
			return true, nil
		}
	}
	if err := scanner.Err(); err != nil {
		log.Printf("%s: %v", op, err)
		return false, fmt.Errorf("%s:%w", op, err)
	}

	return false, nil
}
