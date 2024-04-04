package main

import (
	"github.com/Daanooo/airlabs-fetcher/internal/fetcher"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
)

func main() {
	config.AddDriver(yaml.Driver)

	err := config.LoadFiles("./config.yml")
	if err != nil {
		panic(err)
	}

	params := map[string]any{
		"country_code": "NL",
	}

	fetcher.Fetch("airports", params)
}
