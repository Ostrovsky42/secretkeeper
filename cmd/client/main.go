package main

import (
	"secretkeeper/internal/apps/client/config"
	"secretkeeper/pkg/logger"
)

func main() {
	cfg := config.GetClientConfig()
	log := logger.NewLogger()
	log.Debug().Interface("cfg", cfg).Send()
}
