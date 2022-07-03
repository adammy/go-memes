package meme

import (
	"image"

	"github.com/fogleman/gg"
)

type service struct {
}

// NewService contructs a meme service.
func NewService() (*service, error) {
	return &service{}, nil
}

// CreateMeme creates an image using the provided meme arg.
func (s *service) CreateMeme(meme Meme) (image.Image, error) {
	img, err := gg.LoadImage(meme.ImgPath)
	if err != nil {
		return nil, err
	}

	// dc for "draw context"
	dc := gg.NewContextForImage(img)

	for _, text := range meme.Text {
		if err := drawTextForReal(dc, &text); err != nil {
			return nil, err
		}
	}

	if err := drawWatermark(dc, &meme); err != nil {
		return nil, err
	}

	return dc.Image(), nil
}

func drawTextForReal(dc *gg.Context, text *Text) error {
	if err := loadFont(dc, &text.Font); err != nil {
		return err
	}

	anchorX, anchorY, err := getAnchorCoords(dc, text)
	if err != nil {
		return err
	}

	if err := drawTextStroke(dc, text, anchorX, anchorY); err != nil {
		return err
	}

	if err := drawText(dc, text, anchorX, anchorY); err != nil {
		return err
	}

	return nil
}

func loadFont(dc *gg.Context, font *Font) error {
	fontPath := Fonts[font.Family]
	if fontPath == "" {
		fontPath = DefaultFontPath
	}
	if err := dc.LoadFontFace(fontPath, float64(font.Size)); err != nil {
		return err
	}

	return nil
}

func getAnchorCoords(dc *gg.Context, text *Text) (uint16, uint16, error) {
	lines := len(dc.WordWrap(text.Text, float64(text.Width)))
	x := (text.Width / 2) + text.X
	y := text.Y + (uint16(text.Font.Size/2) * uint16(lines))
	return x, y, nil
}

func drawTextStroke(dc *gg.Context, text *Text, anchorX, anchorY uint16) error {
	if text.Stroke.Enabled {
		dc.SetHexColor(text.Stroke.Color)
		strokeSize := int(text.Stroke.Size)

		for y := -strokeSize; y <= strokeSize; y++ {
			for x := -strokeSize; x <= strokeSize; x++ {
				if x*x+y*y >= strokeSize*strokeSize {
					// give it rounded corners
					continue
				}
				strokeX := anchorX + uint16(x)
				strokeY := anchorY + uint16(y)
				dc.DrawStringWrapped(text.Text, float64(strokeX), float64(strokeY), 0.5, 0.5, float64(text.Width), 1.5, gg.AlignCenter)
			}
		}
	}

	return nil
}

func drawText(dc *gg.Context, text *Text, anchorX, anchorY uint16) error {
	dc.SetHexColor(text.Font.Color)
	dc.DrawStringWrapped(text.Text, float64(anchorX), float64(anchorY), 0.5, 0.5, float64(text.Width), 1.5, gg.AlignCenter)
	return nil
}

func drawWatermark(dc *gg.Context, meme *Meme) error {
	if err := loadFont(dc, &WatermarkFont); err != nil {
		return err
	}
	dc.SetHexColor(WatermarkFont.Color)
	dc.DrawString(WatermarkText, 10, float64(meme.Height-10))
	return nil
}
