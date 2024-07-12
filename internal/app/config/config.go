package config

import (
	"flag"
	"fmt"
	"os"
)

var Config MainConfig

type MainConfig struct {
	BaseServerURL   string
	BaseRedirectURL string
}

func init() {
	flag.StringVar(&Config.BaseServerURL, "a", "localhost:8080", "default host for server")
	flag.StringVar(&Config.BaseRedirectURL, "b", "http://localhost:8080", "default host for server")

	fmt.Println("flags inited")
}

func Parse() {
	flag.Parse()

	if e := os.Getenv("SERVER_ADDRESS"); e != "" {
		Config.BaseServerURL = e
	}
	if e := os.Getenv("BASE_URL"); e != "" {
		Config.BaseRedirectURL = e
	}
}
