package config

import (
	"fmt"
	"net/url"
	cfg "secretkeeper/pkg/configparcer"
)

const Path = "config.yaml"

func GetClientConfig() Config {
	var config Config
	cfg.GetConfig(Path, &config)

	return config
}

type Config struct {
	Postgre Postgre `yaml:"postgre"`
}

type Postgre struct {
	Host           string
	Port           uint16
	User           string
	Password       string
	Database       string
	MinPoolSize    uint8 `yaml:"min_pool_size"`
	MaxPoolSize    uint8 `yaml:"max_pool_size"`
	LogIntervalSec uint8 `yaml:"log_interval_sec"`
	UseSSL         bool
}

func (c Postgre) PostgresDsn() string {
	u := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(c.User, c.Password),
		Host:   fmt.Sprintf("%s:%d", c.Host, c.Port),
		Path:   c.Database,
	}
	if !c.UseSSL {
		u.RawQuery = "sslmode=disable"
	}

	return u.String()
}
