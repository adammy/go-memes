package main

import (
	"strings"

	"github.com/adammy/memepen-services/pkg/font"
	"github.com/adammy/memepen-services/pkg/image"
	"github.com/adammy/memepen-services/pkg/meme"
	"github.com/adammy/memepen-services/pkg/template"
)

func main() {
	fontRepository := font.NewLocalGetter(font.DefaultFontPaths)
	imageRepository := image.NewLocalGetter(image.DefaultImagePaths)
	memeRepository := meme.NewInMemoryRepository()
	templateRepository := template.NewInMemoryRepository(template.DefaultTemplates)
	uploader := image.NewLocalUploader()
	svc := meme.NewService(
		fontRepository,
		imageRepository,
		memeRepository,
		templateRepository,
		uploader,
	)

	_, err := svc.CreateMemeAndUploadFromTemplateID(
		"yall-got-any-more-of-them",
		[]string{
			strings.ToUpper("Y'all Got Any More Of Them"),
			strings.ToUpper("Ape JPEGs"),
		},
	)
	//_, err := svc.CreateMemeAndUploadFromTemplateID(
	//	"two-buttons",
	//	[]string{
	//		"me telling gordo im quitting",
	//		"me just disappearing",
	//		strings.ToUpper("Dav"),
	//	},
	//)
	if err != nil {
		panic(err)
	}
}
