package meme

import (
	"image"

	"github.com/fogleman/gg"
)

// service contains functionality related to creating Meme objects.
type service struct {
	TemplateRepository TemplateRepository
}

// NewService returns a service containing functionality related to creating Meme objects.
func NewService() (*service, error) {
	templateRepository, err := NewInMemoryTemplateRepository()
	if err != nil {
		return nil, err
	}

	return &service{
		TemplateRepository: templateRepository,
	}, nil
}

// CreateMeme creates an image using the provided meme arg.
func (s *service) CreateMeme(templateId string, strs []string) (image.Image, error) {
	template, err := s.TemplateRepository.Get(templateId)
	if err != nil {
		return nil, err
	}

	img, err := gg.LoadImage(template.ImgPath)
	if err != nil {
		return nil, err
	}

	// dc for "draw context"
	dc := gg.NewContextForImage(img)

	for i, text := range template.TextStyle {
		if err := drawTextField(dc, strs[i], &text); err != nil {
			return nil, err
		}
	}

	if err := drawWatermark(dc, template); err != nil {
		return nil, err
	}

	return dc.Image(), nil
}

// drawTextField draws the full text object to the drawing context.
func drawTextField(dc *gg.Context, text string, style *TextStyle) error {
	if err := loadFont(dc, &style.Font); err != nil {
		return err
	}

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

// loadFont loads a font to the drawing context.
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

// getAnchorCoordinates returns the x and y values for the center point of a text field.
func getAnchorCoordinates(dc *gg.Context, text string, style *TextStyle) (uint16, uint16, error) {
	lines := len(dc.WordWrap(text, float64(style.Width)))
	x := (style.Width / 2) + style.X
	y := style.Y + (uint16(style.Font.Size/2) * uint16(lines))
	return x, y, nil
}

// drawTextStroke draws the text stroke/outline to the drawing context.
func drawTextStroke(dc *gg.Context, text string, style *TextStyle, anchorX, anchorY uint16) error {
	if style.Stroke != nil {
		dc.SetHexColor(style.Stroke.Color)
		strokeSize := int(style.Stroke.Size)

		for y := -strokeSize; y <= strokeSize; y++ {
			for x := -strokeSize; x <= strokeSize; x++ {
				if x*x+y*y >= strokeSize*strokeSize {
					// give it rounded corners
					continue
				}
				strokeX := anchorX + uint16(x)
				strokeY := anchorY + uint16(y)
				dc.DrawStringWrapped(text, float64(strokeX), float64(strokeY), 0.5, 0.5, float64(style.Width), 1.5, gg.AlignCenter)
			}
		}
	}

	return nil
}

// drawText draws just the words to the drawing context.
func drawText(dc *gg.Context, text string, style *TextStyle, anchorX, anchorY uint16) error {
	dc.SetHexColor(style.Font.Color)
	dc.DrawStringWrapped(text, float64(anchorX), float64(anchorY), 0.5, 0.5, float64(style.Width), 1.5, gg.AlignCenter)
	return nil
}

// drawWatermark draws a watermark to the drawing context.
func drawWatermark(dc *gg.Context, template *Template) error {
	if err := loadFont(dc, &WatermarkFont); err != nil {
		return err
	}
	dc.SetHexColor(WatermarkFont.Color)
	dc.DrawString(WatermarkText, 10, float64(template.Height-10))
	return nil
}
