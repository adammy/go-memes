package meme

var (
	// Mapping of font name to font path.
	fonts = map[string]string{
		"Impact": "assets/fonts/impact.ttf",
	}

	// Watermark font styling.
	watermarkFont = Font{
		Family: "Impact",
		Size:   15,
		Color:  "#888888",
	}
)
