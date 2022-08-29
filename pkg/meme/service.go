package meme

import (
	"context"
	"math"
	"time"

	"github.com/adammy/memepen-services/pkg/font"
	"github.com/adammy/memepen-services/pkg/img"
	"github.com/adammy/memepen-services/pkg/template"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"github.com/google/uuid"
)

type Service struct {
	fontGetter font.Getter
	imgGetter  img.Getter
	repository Repository
	uploader   img.Uploader
}

func (s *Service) CreateMeme(create CreateMeme) (*Meme, error) {
	img, err := s.imgGetter.Get(create.ImageID)
	if err != nil {
		return nil, err
	}

	dc := gg.NewContextForImage(img)

	for i, text := range create.Text {
		style := create.TextStyles[i]
		font, err := s.fontGetter.Get(string(style.FontFamily))
		if err != nil {
			return nil, err
		}

		if err := drawTextField(dc, text, style, font); err != nil {
			return nil, err
		}
	}

	image := dc.Image()
	id := uuid.NewString()
	if err := s.uploader.UploadPNG(context.Background(), id, image); err != nil {
		return nil, err
	}

	meme := Meme{
		ID:        id,
		CreatedOn: time.Now(),
		URL:       "http://localhost:8080/assets/" + id + ".png",
		ImageID:   create.ImageID,
		Text:      create.Text,
	}

	if err := s.repository.Create(context.Background(), meme); err != nil {
		return nil, err
	}

	return &meme, nil
}

// drawTextField draws the full text object to the drawing context.
func drawTextField(dc *gg.Context, text string, style template.TextStyle, font *truetype.Font) error {
	face := truetype.NewFace(font, &truetype.Options{
		Size: float64(style.FontSize),
	})
	dc.SetFontFace(face)

	anchorX, anchorY, err := getAnchorCoordinates(dc, text, style)
	if err != nil {
		return err
	}

	if err := drawTextStroke(dc, text, style, anchorX, anchorY); err != nil {
		return err
	}

	if err := drawText(dc, text, style, anchorX, anchorY); err != nil {
		return err
	}

	return nil
}

// getAnchorCoordinates returns the x and y values for the center point of a text field.
func getAnchorCoordinates(dc *gg.Context, text string, style template.TextStyle) (int, int, error) {
	lines := len(dc.WordWrap(text, float64(style.Width)))
	x := (style.Width / 2) + style.X
	y := style.Y + ((style.FontSize / 2) * lines)
	return x, y, nil
}

// drawTextStroke draws the text stroke/outline to the drawing context.
func drawTextStroke(dc *gg.Context, text string, style template.TextStyle, anchorX, anchorY int) error {
	if style.StrokeSize != nil && style.StrokeColor != nil {
		dc.SetHexColor(*style.StrokeColor)
		strokeSize := int(*style.StrokeSize)

		for y := -strokeSize; y <= strokeSize; y++ {
			for x := -strokeSize; x <= strokeSize; x++ {
				if x*x+y*y >= strokeSize*strokeSize {
					// give it rounded corners
					continue
				}
				strokeX := anchorX + x
				strokeY := anchorY + y
				if err := rotateText(dc, style, anchorX, anchorY, func() {
					dc.DrawStringWrapped(text, float64(strokeX), float64(strokeY), 0.5, 0.5, float64(style.Width), 1.5, gg.AlignCenter)
				}); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// drawText draws just the words to the drawing context.
func drawText(dc *gg.Context, text string, style template.TextStyle, anchorX, anchorY int) error {
	if err := rotateText(dc, style, anchorX, anchorY, func() {
		dc.SetHexColor(style.FontColor)
		dc.DrawStringWrapped(text, float64(anchorX), float64(anchorY), 0.5, 0.5, float64(style.Width), 1.5, gg.AlignCenter)
	}); err != nil {
		return err
	}

	return nil
}

// rotateText rotates the drawing context and then reverts the rotation after the fn argument is run
func rotateText(dc *gg.Context, style template.TextStyle, anchorX, anchorY int, fn func()) error {
	if style.Rotation != nil {
		radians := gg.Radians(float64(*style.Rotation))
		dc.RotateAbout(radians, float64(anchorX), float64(anchorY))
		if radians >= 0 {
			defer dc.RotateAbout(-radians, float64(anchorX), float64(anchorY))
		} else {
			defer dc.RotateAbout(math.Abs(radians), float64(anchorX), float64(anchorY))
		}
	}
	fn()
	return nil
}
