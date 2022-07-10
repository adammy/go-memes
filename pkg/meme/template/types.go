package template

// Template defines a structure for creating a Meme.
type Template struct {
	// Unique identifier for the template.
	ID string

	// User-friendly name for the template.
	Name string

	// The path to the image. For locally storing, should be "assets/my-image.png". For cloud storing, it'll be some CDN URL.
	ImgPath string

	// The width of the image. We want the precalculated and stored to avoid continually calculating this.
	Width uint16

	// The height of the image. We want the precalculated and stored to avoid continually calculating this.
	Height uint16

	// The text fields for this template.
	TextStyle []TextStyle
}

// TextStyle defines styling and positioning for text.
type TextStyle struct {
	// The x-axis value representing the top-left of a text field.
	X uint16

	// The y-axis value representing the top-left of a text field.
	Y uint16

	// The width of the text field.
	Width uint16

	// Styling for the text.
	Font Font

	// The stroke/outline for the text.
	Stroke *Stroke

	// The rotation for the text.
	Rotation *Rotation
}

// Font defines styling for Text.
type Font struct {
	// The font family (e.g., "Helvetica").
	Family string

	// The font size in point values.
	Size uint8

	// The color of the text
	Color string
}

// Stroke defines the stroke styling for text. It's the border around text.
type Stroke struct {
	// The size of the text stroke.
	Size uint8

	// The color of the text stroke as a hex code, including the prepending pound (#) (e.g., "#FFFFFF").
	Color string
}

// Rotation defines the rotation for text.
type Rotation struct {
	Degrees int16
}
