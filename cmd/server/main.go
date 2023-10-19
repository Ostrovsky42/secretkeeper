package main

import (
	"secretkeeper/internal/apps/server/config"
	"secretkeeper/pkg/logger"
)

func main() {
	cfg := config.GetServerConfig()
	log := logger.NewLogger()
	log.Debug().Interface("cfg", cfg).Send()
}
