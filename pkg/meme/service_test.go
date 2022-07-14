package meme_test

import (
	"testing"

	"github.com/adammy/memepen-services/pkg/meme"
	"github.com/adammy/memepen-services/pkg/meme/font"
	"github.com/adammy/memepen-services/pkg/meme/image"
	"github.com/adammy/memepen-services/pkg/meme/template"
	uploaderPkg "github.com/adammy/memepen-services/pkg/meme/uploader"
	"github.com/stretchr/testify/assert"
)

var (
	fontRepository     = font.NewInMemoryRepository(font.DefaultTestServiceFontPaths)
	imageRepository    = image.NewLocalRepository(image.DefaultTestServiceImagePaths)
	memeRepository     = meme.NewInMemoryRepository()
	templateRepository = template.NewInMemoryRepository(template.DefaultTemplates)
	uploader           = uploaderPkg.NewNoopUploader()
)

func TestNewService(t *testing.T) {
	svc := meme.NewService(
		fontRepository,
		imageRepository,
		memeRepository,
		templateRepository,
		uploader,
	)

	assert.NotNil(t, svc)
}

func TestService_CreateMemeFromTemplateID(t *testing.T) {
	tests := map[string]struct {
		templateID string
		text       []string
		error      bool
	}{
		"valid": {
			templateID: "yall-got-any-more-of-them",
			text:       []string{"test", "test"},
		},
		"invalid": {
			templateID: "not-real",
			text:       []string{"test", "test"},
			error:      true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			svc := meme.NewService(
				fontRepository,
				imageRepository,
				memeRepository,
				templateRepository,
				uploader,
			)
			img, err := svc.CreateMemeFromTemplateID(tc.templateID, tc.text)

			if !tc.error {
				assert.NotNil(t, img)
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
