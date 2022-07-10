package main

import (
	"image/png"
	"os"
	"strings"

	"github.com/adammy/go-memes/pkg/meme"
	"github.com/google/uuid"
)

func main() {
	svc, err := meme.NewService()
	if err != nil {
		panic(err)
	}

	img, err := svc.CreateMeme("yall-got-any-more-of-them", []string{strings.ToUpper("Y'all Got Any More Of Them"), strings.ToUpper("Ape JPEGs")})
	// img, err := svc.CreateMeme("two-buttons", []string{"me telling gordo im quitting", "me just disappearing", strings.ToUpper("Dav")})
	if err != nil {
		panic(err)
	}

	f, err := os.Create(uuid.NewString() + ".png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := png.Encode(f, img); err != nil {
		panic(err)
	}
}
