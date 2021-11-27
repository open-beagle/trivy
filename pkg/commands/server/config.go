package server

import (
	"github.com/urfave/cli/v2"

	"github.com/aquasecurity/trivy/pkg/commands/config"
)

// Config holds the Trivy config
type Config struct {
	config.GlobalConfig
	config.DBConfig
	config.CacheConfig

	Listen      string
	Token       string
	TokenHeader string
}

// NewConfig is the factory method to return config
func NewConfig(c *cli.Context) Config {
	// the error is ignored because logger is unnecessary
	gc, _ := config.NewGlobalConfig(c) // nolint: errcheck
	return Config{
		GlobalConfig: gc,
		DBConfig:     config.NewDBConfig(c),
		CacheConfig:  config.NewCacheConfig(c),

		Listen:      c.String("listen"),
		Token:       c.String("token"),
		TokenHeader: c.String("token-header"),
	}
}

// Init initializes the config
func (c *Config) Init() (err error) {
	if err := c.DBConfig.Init(); err != nil {
		return err
	}
	if err := c.CacheConfig.Init(); err != nil {
		return err
	}

	return nil
}
