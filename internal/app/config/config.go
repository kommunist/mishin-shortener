// Модуль config содержит информацию о настройках приложения.
package config

import (
	"encoding/json"
	"flag"
	"log/slog"
	"os"
)

// Хранит настройки приложения.
type MainConfig struct {
	BaseServerURL   string `json:"base_server_url"`
	BaseRedirectURL string `json:"base_redirect_url"`
	FileStoragePath string `json:"file_storage_path"`
	DatabaseDSN     string `json:"database_dsn"`
	EnableHTTPS     bool   `json:"enable_https"`
}

// Создает структуру харнения с дефолтными значениями.
func MakeConfig() MainConfig {
	config := MainConfig{
		BaseServerURL:   "localhost:8080",
		BaseRedirectURL: "http://localhost:8080",
		FileStoragePath: "",
		DatabaseDSN:     "",
		EnableHTTPS:     false,
	}

	return config
}

// Запускает процесс парсинга флагов и ENV переменных.
func (c *MainConfig) InitConfig() error {
	err := c.getConfigFromJson()
	if err != nil {
		return err
	}
	c.initFlags()
	c.parse()

	return nil
}

func (c *MainConfig) getConfigFromJson() error {
	var jsonConfigPath string
	var data []byte

	flag.StringVar(&jsonConfigPath, "s", "", "json config path")

	// так как у основных настроек ENV выше приоритетом, чем командная строка, то здесь так же
	if e := os.Getenv("CONFIG"); e != "" {
		jsonConfigPath = e
	}

	// открываем файл на чтение
	file, err := os.Open(jsonConfigPath)
	if err != nil {
		slog.Error("Error when open config from json")
		return err
	}
	defer file.Close()

	_, err = file.Read(data)
	if err != nil {
		slog.Error("Error when read json file")
		return err
	}

	err = json.Unmarshal(data, c)
	if err != nil {
		slog.Error("json parsing error")
		return err
	}

	return nil
}

func (c *MainConfig) initFlags() {
	if flag.Lookup("a") == nil {
		flag.StringVar(&c.BaseServerURL, "a", c.BaseServerURL, "default host for server")
		flag.StringVar(&c.BaseRedirectURL, "b", c.BaseRedirectURL, "default host for server")
		flag.StringVar(&c.FileStoragePath, "f", c.FileStoragePath, "file path for file storage")
		flag.StringVar(&c.DatabaseDSN, "d", c.DatabaseDSN, "database DSN")
		flag.BoolVar(&c.EnableHTTPS, "s", c.EnableHTTPS, "database DSN")
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
	if e := os.Getenv("ENABLE_HTTPS"); e != "" {
		if e == "true" || e == "TRUE" {
			c.EnableHTTPS = true
		}
	}
}
