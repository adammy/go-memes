package meme_test

import (
	"testing"

	"github.com/adammy/go-memes/pkg/meme"
	"github.com/stretchr/testify/assert"
)

func TestNewService(t *testing.T) {
	svc, err := meme.NewService("")

	assert.NotNil(t, svc)
	assert.NoError(t, err)
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
			svc, _ := meme.NewService("../../")
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
