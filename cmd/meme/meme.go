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

	// make a meme object
	meme := meme.Meme{
		ImgPath: "assets/yall-got-any-more-of-that.png",
		Width:   600,
		Height:  471,
		Text: []meme.Text{
			{
				X:     10,
				Y:     10,
				Width: 580,
				Text:  strings.ToUpper("Y'all Got Any More Of Them Test Test Test Test"),
				Font: meme.Font{
					Family: "Impact",
					Size:   40,
					Color:  "#FFFFFF",
				},
				Stroke: meme.Stroke{
					Enabled: true,
					Size:    4,
					Color:   "#000000",
				},
			},
			{
				X:     10,
				Y:     421,
				Width: 580,
				Text:  strings.ToUpper("Monkey JPEGs"),
				Font: meme.Font{
					Family: "Impact",
					Size:   40,
					Color:  "#FFFFFF",
				},
				Stroke: meme.Stroke{
					Enabled: true,
					Size:    4,
					Color:   "#000000",
				},
			},
		},
	}

	// create the actual meme using your shitty object
	img, err := svc.CreateMeme(meme)
	if err != nil {
		panic(err)
	}

	// create a new image on the os
	f, err := os.Create("my-garbage-memes/" + uuid.NewString() + ".png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// write meme to image and profit
	if err := png.Encode(f, img); err != nil {
		panic(err)
	}
}
