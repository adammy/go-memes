package memeold_test

import (
	"testing"

	"github.com/adammy/memepen-services/pkg/font"
	"github.com/stretchr/testify/assert"
)

var (
	fontRepository     = font.NewLocalGetter(font.DefaultTestFontPaths)
	imageRepository    = imageold.NewLocalGetter(imageold.DefaultTestImagePaths)
	memeRepository     = memeold.NewInMemoryRepository()
	templateRepository = templateold.NewInMemoryRepository(templateold.DefaultTemplates)
	uploader           = imageold.NewNoopUploader()
)

func TestNewService(t *testing.T) {
	svc := memeold.NewService(
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
			svc := memeold.NewService(
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
