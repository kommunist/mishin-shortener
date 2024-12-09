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
	BaseServerURL   string `json:"server_address"`
	BaseRedirectURL string `json:"base_url"`
	FileStoragePath string `json:"file_storage_path"`
	DatabaseDSN     string `json:"database_dsn"`
	EnableHTTPS     bool   `json:"enable_https"`
	TrustedSubnet   string `json:"trusted_subnet"`
	CertPath        string `json:"cert_path"`
	CertKeyPath     string `json:"cert_key_path"`

	EnableProfile bool
}

// Создает структуру харнения с дефолтными значениями.
func MakeConfig() MainConfig {
	config := MainConfig{
		BaseServerURL:   "localhost:8080",
		BaseRedirectURL: "http://localhost:8080",
		FileStoragePath: "",
		DatabaseDSN:     "",
		EnableHTTPS:     false,
		EnableProfile:   false,
		TrustedSubnet:   "0.0.0.0/32",
		CertPath:        "certs/MyCertificate.crt",
		CertKeyPath:     "certs/MyKey.key",
	}

	return config
}

// Запускает процесс парсинга флагов и ENV переменных.
func (c *MainConfig) InitConfig() error {
	err := c.getConfigFromJSON()
	if err != nil {
		return err
	}
	c.initFlags()
	c.parse()

	return nil
}

func (c *MainConfig) getConfigFromJSON() error {
	var jsonConfigPath string
	var data []byte

	if flag.Lookup("c") == nil {
		flag.StringVar(&jsonConfigPath, "c", "", "json config path")
	}

	// так как у основных настроек ENV выше приоритетом, чем командная строка, то здесь так же
	if e := os.Getenv("CONFIG"); e != "" {
		jsonConfigPath = e
	}

	if jsonConfigPath == "" {
		return nil
	}

	// открываем файл на чтение
	data, err := os.ReadFile(jsonConfigPath)
	if err != nil {
		slog.Error("Error when read json")
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
		flag.StringVar(&c.CertPath, "cert_path", c.CertPath, "path to cert")
		flag.StringVar(&c.CertKeyPath, "cert_key_path", c.CertKeyPath, "path to key of cert")
		flag.StringVar(&c.TrustedSubnet, "t", c.TrustedSubnet, "set trusted subnet")
		flag.BoolVar(&c.EnableProfile, "prof", c.EnableProfile, "start profile server on localhost:6060")
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
	if e := os.Getenv("TRUSTED_SUBNET"); e != "" {
		c.TrustedSubnet = e
	}
	if e := os.Getenv("ENABLE_HTTPS"); e != "" {
		if e == "true" || e == "TRUE" {
			c.EnableHTTPS = true
		}
	}
	if e := os.Getenv("ENABLE_PROFILE"); e != "" {
		if e == "true" || e == "TRUE" {
			c.EnableProfile = true
		}
	}
	if e := os.Getenv("CERT_PATH"); e != "" {
		c.CertPath = e
	}
	if e := os.Getenv("CERT_KEY_PATH"); e != "" {
		c.CertKeyPath = e
	}
}
