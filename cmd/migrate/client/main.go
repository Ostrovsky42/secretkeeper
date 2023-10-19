package main

import (
	"fmt"
	"secretkeeper/internal/apps/client/config"
	"secretkeeper/pkg/migrate"
)

const migrationsPath = "migrations/client"

func main() {
	cfg := config.GetClientConfig().Postgre
	if err := migrate.New(migrationsPath, cfg.PostgresDsn()).Execute(); err != nil {
		fmt.Println(err)
	}
}
