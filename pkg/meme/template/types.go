package template

type RepositoryType string

const (
	InMemory RepositoryType = "InMemory"
)

// Template defines a structure for creating a meme.
type Template struct {
	// ID is a unique identifier for the template.
	ID string

	// Slug is a user-friendly URL slug or identifier.
	Slug string

	// Name is a User-friendly name for the template.
	Name string

	// ImgID is the unique identifier of the image used for the template.
	ImgID string

	// Width is the width of the image. We want the precalculated and stored to avoid continually calculating this.
	Width uint16

	// Height is the height of the image. We want the precalculated and stored to avoid continually calculating this.
	Height uint16

	// TextStyle is the text field styling for this template.
	TextStyle []TextStyle

	// DefaultText is the text used to display meme usage.
	DefaultText []string
}

// TextStyle defines styling and positioning for text.
type TextStyle struct {
	// X is the x-axis value representing the top-left of a text field.
	X uint16

	// Y is the y-axis value representing the top-left of a text field.
	Y uint16

	// Width is the width of the text field.
	Width uint16

	// Font is the font styling for the text.
	Font Font

	// Stroke is the stroke/outline for the text.
	Stroke *Stroke

	// Rotation is the rotation for the text.
	Rotation *Rotation
}

// Font defines styling for text.
type Font struct {
	// Family defines the font family (e.g., "Helvetica").
	Family string

	// Size defines the font size in point values.
	Size uint8

	// Color defines the color of the text in hexadecimal (e.g., "#FFFFFF").
	Color string
}

// Stroke defines the stroke styling for text. It's the border around text.
type Stroke struct {
	// Size defines the size of the text stroke.
	Size uint8

	// Color defines the color of the text stroke in hexadecimal (e.g., "#FFFFFF").
	Color string
}

// Rotation defines the rotation for text.
type Rotation struct {
	// Degrees defines the number of degrees to rotate the text.
	Degrees int16
}
