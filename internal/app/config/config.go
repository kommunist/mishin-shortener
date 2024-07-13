package config

import (
	"flag"
	"log"
	"os"
)

type MainConfig struct {
	BaseServerURL   string
	BaseRedirectURL string
}

func MakeConfig() MainConfig {
	config := MainConfig{
		BaseServerURL:   "localhost:8080",
		BaseRedirectURL: "http://localhost:8080",
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

	log.Printf("flags inited")
}

func (c *MainConfig) Parse() {
	flag.Parse()

	if e := os.Getenv("SERVER_ADDRESS"); e != "" {
		c.BaseServerURL = e
	}
	if e := os.Getenv("BASE_URL"); e != "" {
		c.BaseRedirectURL = e
	}
}
