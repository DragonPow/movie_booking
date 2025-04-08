package config

import (
	"os"

	"github.com/spf13/viper"
)

// Config represents all configuration for the service
type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	JWT      JWTConfig      `mapstructure:"jwt"`
}

// LoadConfig loads configuration from file
func LoadConfig(configFile string) (*Config, error) {
	v := viper.New()
	v.SetConfigFile(configFile)

	// Load config file
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	// Load .env file if it exists
	if _, err := os.Stat(".env"); err == nil {
		v.SetConfigFile(".env")
		v.SetConfigType("env")
		if err := v.MergeInConfig(); err != nil {
			return nil, err
		}
	}

	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
