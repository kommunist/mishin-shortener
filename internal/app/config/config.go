package config

import (
	"flag"
	"fmt"
	"os"
)

var Config MainConfig
type MainConfig struct {
	BaseServerUrl   string
	BaseRedirectUrl string
}

func init() {
	flag.StringVar(&Config.BaseServerUrl, "a", "localhost:8080", "default host for server")
	flag.StringVar(&Config.BaseRedirectUrl, "b", "http://localhost:8080", "default host for server")

	fmt.Println("flags inited")
}

func Parse() {
	flag.Parse()

	if e := os.Getenv("SERVER_ADDRESS"); e != "" {
		Config.BaseServerUrl = e
  }
	if e := os.Getenv("BASE_URL"); e != "" {
		Config.BaseRedirectUrl = e
  }
}
