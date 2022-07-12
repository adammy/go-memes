package meme

import (
	"image"
	"math"

	fontPkg "github.com/adammy/go-memes/pkg/meme/font"
	imagePkg "github.com/adammy/go-memes/pkg/meme/image"
	templatePkg "github.com/adammy/go-memes/pkg/meme/template"
	uploaderPkg "github.com/adammy/go-memes/pkg/meme/uploader"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"github.com/google/uuid"
)

// Service contains functionality related to creating Meme objects.
type Service struct {
	fontRepository     fontPkg.Repository
	imageRepository    imagePkg.Repository
	memeRepository     Repository
	templateRepository templatePkg.Repository
	uploader           uploaderPkg.Uploader
}

// NewService constructs Service.
func NewService(
	fontRepository fontPkg.Repository,
	imageRepository imagePkg.Repository,
	memeRepository Repository,
	templateRepository templatePkg.Repository,
	uploader uploaderPkg.Uploader,
) *Service {
	return &Service{
		fontRepository:     fontRepository,
		imageRepository:    imageRepository,
		memeRepository:     memeRepository,
		templateRepository: templateRepository,
		uploader:           uploader,
	}
}

// CreateMeme creates an image.
func (s *Service) CreateMeme(template *templatePkg.Template, text []string) (image.Image, error) {
	img, err := s.imageRepository.Get(template.ImgID)
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

// CreateMemeAndUpload creates an image and uploads it.
func (s *Service) CreateMemeAndUpload(template *templatePkg.Template, text []string) (*Meme, error) {
	img, err := s.CreateMeme(template, text)
	if err != nil {
		return nil, err
	}

	id := uuid.NewString()
	path := "assets/memes/" + id
	if err := s.uploader.UploadPNG(path, img); err != nil {
		return nil, err
	}

	meme := &Meme{
		ID:         id,
		ImgPath:    "http://localhost:8080/" + path + ".png",
		TemplateID: template.ID,
		Text:       text,
	}
	if err := s.memeRepository.Create(meme); err != nil {
		return nil, err
	}

	return meme, nil
}

// CreateMemeFromTemplateID creates an image using the provided templateID.
func (s *Service) CreateMemeFromTemplateID(templateID string, text []string) (image.Image, error) {
	template, err := s.templateRepository.Get(templateID)
	if err != nil {
		return nil, err
	}

	return s.CreateMeme(template, text)
}

// CreateMemeAndUploadFromTemplateID creates an image using the provided templateID.
func (s *Service) CreateMemeAndUploadFromTemplateID(templateID string, text []string) (*Meme, error) {
	template, err := s.templateRepository.Get(templateID)
	if err != nil {
		return nil, err
	}

	return s.CreateMemeAndUpload(template, text)
}

// drawTextField draws the full text object to the drawing context.
func drawTextField(dc *gg.Context, text string, style *templatePkg.TextStyle, font *truetype.Font) error {
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
func getAnchorCoordinates(dc *gg.Context, text string, style *templatePkg.TextStyle) (uint16, uint16, error) {
	lines := len(dc.WordWrap(text, float64(style.Width)))
	x := (style.Width / 2) + style.X
	y := style.Y + (uint16(style.Font.Size/2) * uint16(lines))
	return x, y, nil
}

// drawTextStroke draws the text stroke/outline to the drawing context.
func drawTextStroke(dc *gg.Context, text string, style *templatePkg.TextStyle, anchorX, anchorY uint16) error {
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
func drawText(dc *gg.Context, text string, style *templatePkg.TextStyle, anchorX, anchorY uint16) error {
	if err := rotateText(dc, style, anchorX, anchorY, func() {
		dc.SetHexColor(style.Font.Color)
		dc.DrawStringWrapped(text, float64(anchorX), float64(anchorY), 0.5, 0.5, float64(style.Width), 1.5, gg.AlignCenter)
	}); err != nil {
		return err
	}

	return nil
}

// rotateText rotates the drawing context and then reverts the rotation after the fn argument is run
func rotateText(dc *gg.Context, style *templatePkg.TextStyle, anchorX, anchorY uint16, fn func()) error {
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
