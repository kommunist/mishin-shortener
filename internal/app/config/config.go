// Модуль config содержит информацию о настройках приложения.
package config

import (
	"flag"
	"log/slog"
	"os"
)

// Хранит настройки приложения.
type MainConfig struct {
	BaseServerURL   string
	BaseRedirectURL string
	FileStoragePath string
	DatabaseDSN     string
}

// Создает структуру харнения с дефолтными значениями.
func MakeConfig() MainConfig {
	config := MainConfig{
		BaseServerURL:   "localhost:8080",
		BaseRedirectURL: "http://localhost:8080",
		FileStoragePath: "",
		DatabaseDSN:     "",
	}

	return config
}

// Запускает процесс парсинга флагов и ENV переменных.
func (c *MainConfig) InitConfig() {
	c.initFlags()
	c.parse()
}

func (c *MainConfig) initFlags() {
	if flag.Lookup("a") == nil {
		flag.StringVar(&c.BaseServerURL, "a", "localhost:8080", "default host for server")
		flag.StringVar(&c.BaseRedirectURL, "b", "http://localhost:8080", "default host for server")
		flag.StringVar(&c.FileStoragePath, "f", "", "file path for file storage")
		flag.StringVar(&c.DatabaseDSN, "d", "", "database DSN")
		slog.Info("flags inited")
	}
}

func (c *MainConfig) parse() {
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
}
