package config

import (
	"crypto/ed25519"
	"log"
	"strings"

	"github.com/mohammadanang/shopping-cart/pkg/key"
	"github.com/spf13/viper"
)

type Env struct {
	AppName        string `mapstructure:"APP_NAME"`
	AppEnv         string `mapstructure:"APP_ENV"`
	Port           string `mapstructure:"PORT"`
	AllowedOrigins string `mapstructure:"ALLOWED_ORIGINS"`
	AllowedMethods string `mapstructure:"ALLOWED_METHODS"`

	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     int    `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	SSLMode    string `mapstructure:"DB_SSLMODE"`

	XenditBaseUrl      string `mapstructure:"XENDIT_BASE_URL"`
	XenditSecretKey    string `mapstructure:"XENDIT_SECRET_KEY"`
	XenditWebhookToken string `mapstructure:"XENDIT_WEBHOOK_TOKEN"`
}

type Config struct {
	Env        *Env
	PrivateKey ed25519.PrivateKey
	PublicKey  ed25519.PublicKey
}

func NewConfig(path string) *Config {
	env := loadEnv(path)
	keyLoader := key.NewKeyLoader(path)

	return &Config{
		Env:        env,
		PrivateKey: keyLoader.LoadPrivateKey(),
		PublicKey:  keyLoader.LoadPublicKey(),
	}
}

func loadEnv(path string) *Env {
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
		log.Fatalf("No config file found at %s, relying only on environment variables", path)
	}

	var env Env
	// Unmarshal into struct
	if err := v.Unmarshal(&env); err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	return &env
}
