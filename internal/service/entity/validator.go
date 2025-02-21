package validator

import "github.com/go-playground/validator/v10"

func HostnameValidate(hostname string) error {
	val := validator.New(validator.WithRequiredStructEnabled())

	err := val.Var(hostname, "required,hostname_rfc1123")

	if err != nil {
		return err
	}
	return nil
}

func IPValidate(dnsServer string) error {
	val := validator.New(validator.WithRequiredStructEnabled())

	err := val.Var(dnsServer, "required,ipv4")

	if err != nil {
		return err
	}
	return nil
}
