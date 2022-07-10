package main

import (
	"github.com/adammy/go-memes/pkg/meme/server"
)

func main() {
	server, err := server.NewGinServer()
	if err != nil {
		panic(err)
	}

	if err := server.Start(); err != nil {
		panic(err)
	}
}
