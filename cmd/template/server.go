package main

import (
	"github.com/adammy/memepen-services/pkg/template/server"
)

func main() {
	config, err := server.LoadConfig("./configs", "local")
	if err != nil {
		panic(err)
	}

	s, err := server.NewGinServer(config)
	if err != nil {
		panic(err)
	}

	if err := s.Start(); err != nil {
		panic(err)
	}
}
