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
	config, err := server.LoadConfig("./configs", "local")
	if err != nil {
		panic(err)
	}

	fontRepository := font.NewRepository(config.FontRepositoryType, font.DefaultFontPaths)
	imageRepository := image.NewRepository(config.ImageRepositoryType, image.DefaultImagePaths)
	memeRepository := meme.NewRepository(config.MemeRepositoryType)
	templateRepository := template.NewRepository(config.TemplateRepositoryType, template.DefaultTemplates)
	uploader := uploaderPkg.NewUploader(config.UploaderType)
	service := meme.NewService(
		fontRepository,
		imageRepository,
		memeRepository,
		templateRepository,
		uploader,
	)

	s, err := server.NewGinServer(config, service)
	if err != nil {
		panic(err)
	}

	if err := s.Start(); err != nil {
		panic(err)
	}
}
