package main

import (
	"fmt"
	"hostsetup-service/internal/service/config"
)

func main() {
	// TODO: read config
	cfg := config.MustLoadConfig()
	fmt.Println(cfg)

	// TODO: start app

	// TODO: graceful shutdown
}
