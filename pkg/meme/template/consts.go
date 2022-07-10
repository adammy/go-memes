package template

var (
	DefaultTemplates = map[string]Template{
		"yall-got-any-more-of-them": {
			ImgPath: "assets/templates/yall-got-any-more-of-that.png",
			Width:   600,
			Height:  471,
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
			ImgPath: "assets/templates/two-buttons.png",
			Width:   500,
			Height:  756,
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
