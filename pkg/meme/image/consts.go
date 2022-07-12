package image

var (
	// DefaultImagePaths defines the images available for meme templates.
	DefaultImagePaths = map[string]string{
		"yall-got-any-more-of-them": "assets/templates/yall-got-any-more-of-that.png",
		"two-buttons":               "assets/templates/two-buttons.png",
	}

	// DefaultTestImagePaths defines the images available for meme templates.
	DefaultTestImagePaths = map[string]string{
		"yall-got-any-more-of-them": "../../../assets/templates/yall-got-any-more-of-that.png",
		"two-buttons":               "../../../assets/templates/two-buttons.png",
	}

	// DefaultTestServiceImagePaths defines the images available for meme templates.
	DefaultTestServiceImagePaths = map[string]string{
		"yall-got-any-more-of-them": "../../assets/templates/yall-got-any-more-of-that.png",
		"two-buttons":               "../../assets/templates/two-buttons.png",
	}
)
