package main

import (
	"image/png"
	"os"
	"strings"

	"github.com/adammy/go-memes/pkg/meme"
	"github.com/adammy/go-memes/pkg/meme/font"
	imagePkg "github.com/adammy/go-memes/pkg/meme/image"
	"github.com/adammy/go-memes/pkg/meme/template"
	uploaderPkg "github.com/adammy/go-memes/pkg/meme/uploader"
	"github.com/google/uuid"
)

func main() {
	fontRepository := font.NewInMemoryRepository(font.DefaultFontPaths)
	imageRepository := imagePkg.NewLocalRepository(imagePkg.DefaultImagePaths)
	memeRepository := meme.NewInMemoryRepository()
	templateRepository := template.NewInMemoryRepository(template.DefaultTemplates)
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
