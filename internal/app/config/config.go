package config

import (
	"flag"
	"log/slog"
	"os"
)

type MainConfig struct {
	BaseServerURL   string
	BaseRedirectURL string
	FileStoragePath string
	DatabaseDSN     string
	SecureKey       string
}

func MakeConfig() MainConfig {
	config := MainConfig{
		BaseServerURL:   "localhost:8080",
		BaseRedirectURL: "http://localhost:8080",
		FileStoragePath: "",
		DatabaseDSN:     "",
		SecureKey:       "0123456789123456",
	}

	return config
}

func (c *MainConfig) InitConfig() {
	c.InitFlags()
	c.Parse()
}

func (c *MainConfig) InitFlags() {
	flag.StringVar(&c.BaseServerURL, "a", "localhost:8080", "default host for server")
	flag.StringVar(&c.BaseRedirectURL, "b", "http://localhost:8080", "default host for server")
	flag.StringVar(&c.FileStoragePath, "f", "", "file path for file storage")
	flag.StringVar(&c.DatabaseDSN, "d", "", "database DSN")
	flag.StringVar(&c.SecureKey, "secure", "12345678901234567890123456789012", "secure key for encrypt")

	slog.Info("flags inited")
}

func (c *MainConfig) Parse() {
	flag.Parse()

	if e := os.Getenv("SERVER_ADDRESS"); e != "" {
		c.BaseServerURL = e
	}
	if e := os.Getenv("BASE_URL"); e != "" {
		c.BaseRedirectURL = e
	}
	if e := os.Getenv("FILE_STORAGE_PATH"); e != "" {
		c.FileStoragePath = e
	}
	if e := os.Getenv("DATABASE_DSN"); e != "" {
		c.DatabaseDSN = e
	}
	if e := os.Getenv("SECURE_KEY"); e != "" {
		c.SecureKey = e
	}
}
