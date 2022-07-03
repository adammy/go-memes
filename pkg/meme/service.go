package meme

import (
	"image"

	"github.com/fogleman/gg"
	"github.com/g4s8/hexcolor"
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

	dc := gg.NewContext(int(meme.Width), int(meme.Height))
	dc.DrawImage(img, 0, 0)

	for _, text := range meme.Text {
		if err := s.addText(dc, text); err != nil {
			return nil, err
		}
	}

	return dc.Image(), nil
}

func (s *service) addText(dc *gg.Context, text Text) error {
	fontPath := Fonts[text.Font.Family]
	if fontPath == "" {
		fontPath = DefaultFontPath
	}
	if err := dc.LoadFontFace(fontPath, float64(text.Font.Size)); err != nil {
		return err
	}

	color, err := hexcolor.Parse(text.Font.Color)
	if err != nil {
		return err
	}
	dc.SetColor(color)

	x := (text.Width / 2) + text.X
	y := text.Y + uint16(text.Font.Size/2)
	dc.DrawStringWrapped(text.Text, float64(x), float64(y), 0.5, 0.5, float64(text.Width), 1.5, gg.AlignCenter)

	return nil
}
