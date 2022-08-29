package main

import (
	"github.com/adammy/memepen-services/pkg/httpapi"
	"github.com/adammy/memepen-services/pkg/img"
)

func main() {
	config, err := httpapi.LoadConfig[img.Config](httpapi.ConfigPath, "img")
	if err != nil {
		panic(err)
	}

	server := img.NewServer(config)
	if err := server.Start(); err != nil {
		panic(err)
	}
}
