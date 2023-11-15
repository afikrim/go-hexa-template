package config

import (
	"fmt"
	"os"
	"path"

	"github.com/afikrim/go-hexa-template/pkg/envloader"
)

type Config struct {
	// Application Configurations
	Application Application
	// Database Configurations
	Database Database
}

// ConfigOption is a function to set config options
type ConfigOption func(*Config) error

// New is a function to create new config
func New(opts ...ConfigOption) *Config {
	cfg := &Config{}

	for _, opt := range opts {
		err := opt(cfg)
		if err != nil {
			fmt.Println("error while setting config", err)
		}
	}

	return cfg
}

// WithEtcd is a function to set config from etcd
func WithEtcd(prefix string, endpoints []string) ConfigOption {
	return func(cfg *Config) error {
		return envloader.Load(cfg, envloader.WithEtcd(prefix, endpoints))
	}
}

// WithFile is a function to set config file
func WithFile(file string) ConfigOption {
	return func(cfg *Config) error {
		wd, err := os.Getwd()
		if err != nil {
			return err
		}

		return envloader.Load(cfg, envloader.WithFile(path.Join(wd, file)))
	}
}

// WithEnv is a function to set config from environment
func WithEnv() ConfigOption {
	return func(cfg *Config) error {
		return envloader.Load(cfg)
	}
}
