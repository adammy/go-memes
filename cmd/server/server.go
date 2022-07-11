package main

import (
	"github.com/adammy/go-memes/pkg/meme/server"
)

func main() {
	s, err := server.NewGinServer()
	if err != nil {
		panic(err)
	}

	if err := s.Start(); err != nil {
		panic(err)
	}
}
