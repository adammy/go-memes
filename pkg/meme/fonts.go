package meme

var (
	// Mapping of font name to font path.
	Fonts = map[string]string{
		"Impact": "assets/fonts/impact.ttf",
	}

	// Watermark font styling.
	WatermarkFont = Font{
		Family: "Impact",
		Size:   15,
		Color:  "#888888",
	}
)
