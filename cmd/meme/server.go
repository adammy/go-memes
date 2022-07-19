package main

import (
	"github.com/adammy/memepen-services/pkg/font"
	"github.com/adammy/memepen-services/pkg/image"
	"github.com/adammy/memepen-services/pkg/meme"
	"github.com/adammy/memepen-services/pkg/template"
)

func main() {
	config, err := meme.LoadConfig("./configs", "local")
	if err != nil {
		panic(err)
	}

	fontRepository := font.NewGetter(config.FontRepositoryType, font.DefaultFontPaths)
	imageRepository := image.NewGetter(config.ImageRepositoryType, image.DefaultImagePaths)
	memeRepository := meme.NewRepository(config.MemeRepositoryType)
	templateRepository := template.NewRepository(config.TemplateRepositoryType, template.DefaultTemplates)
	uploader := image.NewUploader(config.UploaderType)
	service := meme.NewService(
		fontRepository,
		imageRepository,
		memeRepository,
		templateRepository,
		uploader,
	)

	s, err := meme.NewGinServer(config, service)
	if err != nil {
		panic(err)
	}

	if err := s.Start(); err != nil {
		panic(err)
	}
}
