package main

import (
	"flag"

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

	fileFlag := ""
	flag.StringVar(&fileFlag, "file", "", "")
	flag.Parse()

	airports := fetcher.Fetch("airports", flag.Args())

	if err := fetcher.MakeExport(airports, fileFlag); err != nil {
		panic(err)
	}
}
