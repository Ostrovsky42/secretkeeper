package main

import (
	"fmt"
	"secretkeeper/internal/apps/server/config"
	"secretkeeper/pkg/migrate"
)

const migrationsPath = "migrations/server"

func main() {
	cfg := config.GetServerConfig().Postgre

	if err := migrate.New(migrationsPath, cfg.PostgresDsn()).Execute(); err != nil {
		fmt.Println(err)
	}
}
