package main

import (
	"github.com/adammy/memepen-services/pkg/template"
)

func main() {
	config, err := template.LoadConfig("./configs", "local")
	if err != nil {
		panic(err)
	}

	s, err := template.NewGinServer(config)
	if err != nil {
		panic(err)
	}

	if err := s.Start(); err != nil {
		panic(err)
	}
}
