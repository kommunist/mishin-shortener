package config

import (
	"flag"
	"fmt"
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
}
