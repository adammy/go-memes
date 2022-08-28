package main

import (
	"github.com/adammy/memepen-services/pkg/httpapi"
	"github.com/adammy/memepen-services/pkg/template"
)

func main() {
	config, err := httpapi.LoadConfig[template.Config](httpapi.ConfigPath, "template")
	if err != nil {
		panic(err)
	}

	server := template.NewServer(config)
	if err := server.Start(); err != nil {
		panic(err)
	}
}
