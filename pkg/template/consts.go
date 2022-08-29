package template

import (
	"time"

	"github.com/adammy/memepen-services/pkg/color"
	"github.com/adammy/memepen-services/pkg/pointer"
)

const (
	// InMemoryRepository denotes to use an in-memory map for templates.
	InMemoryRepository RepositoryType = "InMemoryRepository"

	// PostgresRepository denotes to use a postgres db for templates.
	PostgresRepository RepositoryType = "PostgresRepository"
)

var (
	DefaultTemplates = map[string]*Template{
		"yall-got-any-more-of-them": {
			ID:        "yall-got-any-more-of-them",
			Name:      "Y'all Got Any More of Them",
			CreatedOn: time.Now(),
			ImageID:   "yall-got-any-more-of-them",
			TextStyles: []TextStyle{
				{
					X:           10,
					Y:           10,
					Width:       580,
					FontFamily:  Impact,
					FontSize:    40,
					FontColor:   color.White,
					StrokeSize:  pointer.GetIntP(4),
					StrokeColor: pointer.GetStringP(color.Black),
				},
				{
					X:           10,
					Y:           421,
					Width:       580,
					FontFamily:  Impact,
					FontSize:    40,
					FontColor:   color.White,
					StrokeSize:  pointer.GetIntP(4),
					StrokeColor: pointer.GetStringP(color.Black),
				},
			},
			DefaultText: []string{"Y'all Got Any More of Them", "Something"},
		},
		"two-buttons": {
			ID:        "two-buttons",
			Name:      "Two Buttons",
			CreatedOn: time.Now(),
			ImageID:   "two-buttons",
			TextStyles: []TextStyle{
				{
					X:          80,
					Y:          110,
					Width:      100,
					FontFamily: Arial,
					FontSize:   20,
					FontColor:  color.Black,
					Rotation:   pointer.GetIntP(-10),
				},
				{
					X:          245,
					Y:          80,
					Width:      100,
					FontFamily: Arial,
					FontSize:   20,
					FontColor:  color.Black,
					Rotation:   pointer.GetIntP(-10),
				},
				{
					X:           20,
					Y:           675,
					Width:       460,
					FontFamily:  Impact,
					FontSize:    40,
					FontColor:   color.Black,
					StrokeSize:  pointer.GetIntP(4),
					StrokeColor: pointer.GetStringP(color.Black),
				},
			},
			DefaultText: []string{"Option 1", "Option 2", "Person Deciding"},
		},
	}
)
