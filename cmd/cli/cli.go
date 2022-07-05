package main

import (
	"image/png"
	"os"
	"strings"

	"github.com/adammy/go-memes/pkg/meme"
	"github.com/google/uuid"
)

func main() {
	// constuct the service
	svc, err := meme.NewService()
	if err != nil {
		panic(err)
	}

	// create the actual meme using your shitty object
	img, err := svc.CreateMeme("123", []string{strings.ToUpper("Y'all Got Any More Of Them"), strings.ToUpper("Ape JPEGs")})
	if err != nil {
		panic(err)
	}

	// create a new image on the os
	f, err := os.Create(uuid.NewString() + ".png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// write meme to image and profit
	if err := png.Encode(f, img); err != nil {
		panic(err)
	}
}
