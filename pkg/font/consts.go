package font

const (
	// LocalGetter denotes to use a local getter for fonts.
	LocalGetter GetterType = "LocalGetter"
)

var (
	// DefaultFontPaths defines the fonts and their paths.
	DefaultFontPaths = map[string]string{
		"Arial":  "assets/fonts/arial.ttf",
		"Impact": "assets/fonts/impact.ttf",
	}

	// DefaultTestFontPaths defines the fonts and their paths relative to test files.
	DefaultTestFontPaths = map[string]string{
		"Arial":  "../../assets/fonts/arial.ttf",
		"Impact": "../../assets/fonts/impact.ttf",
	}
)
