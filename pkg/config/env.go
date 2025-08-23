package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	AppName string `mapstructure:"APP_NAME"`
	AppEnv  string `mapstructure:"APP_ENV"`
	Port    string `mapstructure:"PORT"`

	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     int    `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	SSLMode    string `mapstructure:"DB_SSLMODE"`
}

var C Config

// LoadConfig loads environment variables into Config struct
func LoadConfig(path string) *Config {
	v := viper.New()

	// Look for .env file
	v.SetConfigName("app") // app.env or app.yaml
	v.SetConfigType("env") // or "yaml"
	v.AddConfigPath(path)
	v.AutomaticEnv()

	// Replace dots with underscores if needed
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Read file if exists
	if err := v.ReadInConfig(); err != nil {
		log.Printf("No config file found at %s, relying only on environment variables", path)
	}

	// Unmarshal into struct
	if err := v.Unmarshal(&C); err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	return &C
}
