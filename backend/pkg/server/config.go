package server

import "fmt"

// Config is the configuration for the server.
type Config struct {
	Port         int      `yaml:"port"`
	UseTLS       bool     `yaml:"use_tls"`
	CrtPath      string   `yaml:"tls_crt"`
	KeyPath      string   `yaml:"tls_key"`
	CORS         bool     `yaml:"cors"`
	AllowOrigins []string `yaml:"allow_origins"`
}

func (c *Config) Addr() string {
	if c == nil {
		return ""
	}
	return fmt.Sprintf(":%d", c.Port)
}

// DefaultConfig returns the default configuration for the server.
func DefaultConfig() *Config {
	return &Config{
		Port:         8080,
		UseTLS:       false,
		CORS:         false,
		AllowOrigins: []string{},
	}
}
