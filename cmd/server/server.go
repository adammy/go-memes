package main

import (
	"github.com/adammy/go-memes/pkg/meme"
	"github.com/adammy/go-memes/pkg/meme/font"
	"github.com/adammy/go-memes/pkg/meme/image"
	"github.com/adammy/go-memes/pkg/meme/server"
	"github.com/adammy/go-memes/pkg/meme/template"
	uploaderPkg "github.com/adammy/go-memes/pkg/meme/uploader"
)

func main() {
	fontRepository := font.NewInMemoryRepository(font.DefaultFontPaths)
	imageRepository := image.NewLocalRepository(image.DefaultImagePaths)
	memeRepository := meme.NewInMemoryRepository()
	templateRepository := template.NewInMemoryRepository(template.DefaultTemplates)
	uploader := uploaderPkg.NewLocalUploader()
	service := meme.NewService(
		fontRepository,
		imageRepository,
		memeRepository,
		templateRepository,
		uploader,
	)

	s, err := server.NewGinServer(service)
	if err != nil {
		panic(err)
	}

	if err := s.Start(); err != nil {
		panic(err)
	}
}
