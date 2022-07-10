package meme

import (
	"image"
	"math"

	"github.com/adammy/go-memes/pkg/meme/font"
	"github.com/adammy/go-memes/pkg/meme/template"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
)

// service contains functionality related to creating Meme objects.
type service struct {
	templateRepository template.Repository
	fontRepository     font.Repository
}

// NewService constructs service.
func NewService() (*service, error) {
	templateRepository, err := template.NewInMemoryRepository(nil)
	if err != nil {
		return nil, err
	}

	fontRepository, err := font.NewInMemoryRepository(nil)
	if err != nil {
		return nil, err
	}

	return &service{
		templateRepository: templateRepository,
		fontRepository:     fontRepository,
	}, nil
}

// CreateMeme creates an image using the provided meme arg.
func (s *service) CreateMeme(templateId string, text []string) (image.Image, error) {
	template, err := s.templateRepository.Get(templateId)
	if err != nil {
		return nil, err
	}

	img, err := gg.LoadImage(template.ImgPath)
	if err != nil {
		return nil, err
	}

	dc := gg.NewContextForImage(img)

	for i, style := range template.TextStyle {
		font, err := s.fontRepository.Get(style.Font.Family)
		if err != nil {
			return nil, err
		}

		if err := drawTextField(dc, text[i], &style, font); err != nil {
			return nil, err
		}
	}

	return dc.Image(), nil
}

// drawTextField draws the full text object to the drawing context.
func drawTextField(dc *gg.Context, text string, style *template.TextStyle, font *truetype.Font) error {
	face := truetype.NewFace(font, &truetype.Options{
		Size: float64(style.Font.Size),
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
func getAnchorCoordinates(dc *gg.Context, text string, style *template.TextStyle) (uint16, uint16, error) {
	lines := len(dc.WordWrap(text, float64(style.Width)))
	x := (style.Width / 2) + style.X
	y := style.Y + (uint16(style.Font.Size/2) * uint16(lines))
	return x, y, nil
}

// drawTextStroke draws the text stroke/outline to the drawing context.
func drawTextStroke(dc *gg.Context, text string, style *template.TextStyle, anchorX, anchorY uint16) error {
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
func drawText(dc *gg.Context, text string, style *template.TextStyle, anchorX, anchorY uint16) error {
	if err := rotateText(dc, style, anchorX, anchorY, func() {
		dc.SetHexColor(style.Font.Color)
		dc.DrawStringWrapped(text, float64(anchorX), float64(anchorY), 0.5, 0.5, float64(style.Width), 1.5, gg.AlignCenter)
	}); err != nil {
		return err
	}

	return nil
}

// rotateText rotates the drawing context and then reverts the rotation after the fn argument is run
func rotateText(dc *gg.Context, style *template.TextStyle, anchorX, anchorY uint16, fn func()) error {
	if style.Rotation != nil {
		radians := gg.Radians(float64(style.Rotation.Degrees))
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
