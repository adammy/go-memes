package main

import (
	"github.com/adammy/memepen-services/pkg/meme"
	"github.com/adammy/memepen-services/pkg/httpapi"
)

func main() {
	config, err := httpapi.LoadConfig[meme.Config](httpapi.ConfigPath, "meme")
	if err != nil {
		panic(err)
	}

	server := meme.NewServer(config)
	if err := server.Start(); err != nil {
		panic(err)
	}
}
