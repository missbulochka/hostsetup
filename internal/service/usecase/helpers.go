package setupping

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func verifyHostname(hostname string) error {
	currentHostname, err := os.Hostname()
	if err != nil {
		return err
	}
	if currentHostname != hostname {
		return fmt.Errorf("%s", "failed to change hostname")
	}

	return nil
}

func dnsServerExists(file *os.File, resolverString string) (bool, error) {
	resolverString = strings.Trim(resolverString, " ")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " ")
		if line == resolverString {
			return true, nil
		}
	}
	if err := scanner.Err(); err != nil {
		return false, err
	}

	return false, nil
}

func resolvFileBackup(originFile *os.File) (*os.File, error) {
	fileBackup, err := os.OpenFile(pwdToResolvConfBackup, os.O_CREATE|os.O_WRONLY, 0222)
	if err != nil {
		return nil, err
	}

	if _, err := io.Copy(fileBackup, originFile); err != nil {
		fileBackup.Close()
		return nil, err
	}

	if _, err := originFile.Seek(0, 0); err != nil {
		fileBackup.Close()
		return nil, err
	}

	return fileBackup, nil
}

func removingStringInData(file *os.File, resolverStringToDelete string) (string, error) {
	exist := false
	var newFile string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if resolverStringToDelete != strings.TrimSpace(line) {
			newFile = newFile + line + "\n"
			continue
		}
		exist = true
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	if !exist {
		return "", fmt.Errorf("%s", "dns server doesn't exist")
	}

	return newFile, nil
}
