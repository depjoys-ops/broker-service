package main

import (
	"github.com/depjoys-ops/broker-service/internal/api"
	"github.com/depjoys-ops/broker-service/internal/config"
)

func main() {
	cfg := config.Load()
	api.Run(cfg)
}
