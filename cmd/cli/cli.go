package main

import (
	"image/png"
	"os"
	"strings"

	"github.com/adammy/memepen-services/pkg/meme"
	"github.com/adammy/memepen-services/pkg/meme/font"
	imagePkg "github.com/adammy/memepen-services/pkg/meme/image"
	"github.com/adammy/memepen-services/pkg/template"
	"github.com/adammy/memepen-services/pkg/template/repository"
	uploaderPkg "github.com/adammy/memepen-services/pkg/uploader"
	"github.com/google/uuid"
)

func main() {
	fontRepository := font.NewInMemoryRepository(font.DefaultFontPaths)
	imageRepository := imagePkg.NewLocalRepository(imagePkg.DefaultImagePaths)
	memeRepository := meme.NewInMemoryRepository()
	templateRepository := repository.NewInMemoryRepository(template.DefaultTemplates)
	uploader := uploaderPkg.NewLocalUploader()
	svc := meme.NewService(
		fontRepository,
		imageRepository,
		memeRepository,
		templateRepository,
		uploader,
	)

	img, err := svc.CreateMemeFromTemplateID("yall-got-any-more-of-them", []string{strings.ToUpper("Y'all Got Any More Of Them"), strings.ToUpper("Ape JPEGs")})
	// img, err := svc.CreateMemeFromTemplateID("two-buttons", []string{"me telling gordo im quitting", "me just disappearing", strings.ToUpper("Dav")})
	if err != nil {
		panic(err)
	}

	file, err := os.Create(uuid.NewString() + ".png")
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(file)

	if err := png.Encode(file, img); err != nil {
		panic(err)
	}
}
