package config

import "github.com/kelseyhightower/envconfig"

type GRPCConfig struct {
	GRPCServer string `envconfig:"HOSTSETUP_SERVER" default:"0.0.0.0"`
	GRPCPort   string `envconfig:"HOSTSETUP_PORT" default:"8081"`
}

// MustLoadConfig reads config from env and init *Config value
func MustLoadConfig() *GRPCConfig {
	var cfg = new(GRPCConfig)
	if err := envconfig.Process("", cfg); err != nil {
		panic("hostsetup: failed to read config: " + err.Error())
	}

	return cfg
}
