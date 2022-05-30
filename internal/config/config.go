package config

import (
	"context"

	envconfig "github.com/sethvargo/go-envconfig"

	"github.com/trevatk/common/database"
	"github.com/trevatk/common/server"
)

// Config class to hold all service configs
type Config struct {
	Database database.Config
	Server   server.Server
}

// ProvideConfig create new config instance
func ProvideConfig() *Config {
	return &Config{}
}

// InvokeConfig
func InvokeConfig(cfg *Config) error {
	return envconfig.ProcessWith(context.TODO(), cfg, envconfig.OsLookuper())
}
