package template

var (
	// DefaultTemplates defines the templates available for memes.
	DefaultTemplates = map[string]*Template{
		"yall-got-any-more-of-them": {
			ID:          "yall-got-any-more-of-them",
			Slug:        "yall-got-any-more-of-them",
			Name:        "Y'all Got Any More of Them",
			ImgID:       "yall-got-any-more-of-them",
			Width:       600,
			Height:      471,
			DefaultText: []string{"Y'all Got Any More of Them", "Something"},
			TextStyle: []TextStyle{
				{
					X:     10,
					Y:     10,
					Width: 580,
					Font: Font{
						Family: "Impact",
						Size:   40,
						Color:  "#FFFFFF",
					},
					Stroke: &Stroke{
						Size:  4,
						Color: "#000000",
					},
				},
				{
					X:     10,
					Y:     421,
					Width: 580,
					Font: Font{
						Family: "Impact",
						Size:   40,
						Color:  "#FFFFFF",
					},
					Stroke: &Stroke{
						Size:  4,
						Color: "#000000",
					},
				},
			},
		},
		"two-buttons": {
			ID:          "two-buttons",
			Slug:        "yall-got-any-more-of-them",
			Name:        "Two Buttons",
			ImgID:       "two-buttons",
			Width:       500,
			Height:      756,
			DefaultText: []string{"Option 1", "Option 2", "Person Deciding"},
			TextStyle: []TextStyle{
				{
					X:     80,
					Y:     110,
					Width: 100,
					Font: Font{
						Family: "Arial",
						Size:   20,
						Color:  "#000000",
					},
					Rotation: &Rotation{
						Degrees: -10,
					},
				},
				{
					X:     245,
					Y:     80,
					Width: 100,
					Font: Font{
						Family: "Arial",
						Size:   20,
						Color:  "#000000",
					},
					Rotation: &Rotation{
						Degrees: -10,
					},
				},
				{
					X:     20,
					Y:     675,
					Width: 460,
					Font: Font{
						Family: "Impact",
						Size:   40,
						Color:  "#FFFFFF",
					},
					Stroke: &Stroke{
						Size:  4,
						Color: "#000000",
					},
				},
			},
		},
	}
)
