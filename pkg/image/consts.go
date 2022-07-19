package image

const (
	// LocalGetter denotes to use the os filesystem to get images.
	LocalGetter GetterType = "Local"

	// LocalUploader denotes to upload items to the os filesystem.
	LocalUploader UploaderType = "Local"

	// NoopUploader denotes to not upload items and fake success.
	NoopUploader UploaderType = "Noop"
)

var (
	// DefaultImagePaths defines the images available for meme templates.
	DefaultImagePaths = map[string]string{
		"yall-got-any-more-of-them": "assets/templates/yall-got-any-more-of-that.png",
		"two-buttons":               "assets/templates/two-buttons.png",
	}

	// DefaultTestImagePaths defines the images available for meme templates.
	DefaultTestImagePaths = map[string]string{
		"yall-got-any-more-of-them": "../../assets/templates/yall-got-any-more-of-that.png",
		"two-buttons":               "../../assets/templates/two-buttons.png",
	}
)
