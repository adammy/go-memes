package template

import (
	"time"

	"github.com/adammy/memepen-services/pkg/httpapi"
	"github.com/adammy/memepen-services/pkg/image"
)

// RepositoryType denotes the type of Repository to use for templates.
type RepositoryType string

// Template defines a structure for creating a meme.
type Template struct {
	// ID is a unique identifier for the template.
	ID string `json:"id"`

	// Slug is a user-friendly URL slug or identifier.
	Slug string `json:"slug"`

	// Name is a User-friendly name for the template.
	Name string `json:"name"`

	// NSFW defines if the template is inappropriate for work or not.
	NSFW bool `json:"nsfw"`

	// CreatedOn defines the date/time the template was created.
	CreatedOn time.Time `json:"createdOn"`

	// UserID defines the user who created the template.
	UserID string `json:"userId"`

	// ImageID define the base image used for the template.
	Image Image `json:"image"`

	// TextStyles is the text field styling for this template.
	TextStyles []TextStyle `json:"textStyles"`

	// DefaultText is the text used to display meme usage.
	DefaultText []string `json:"defaultText"`
}

// Image denotes a base image for a template.
type Image struct {
	// ID is the unique identifier of the image used for the template.
	ID string `json:"id"`

	// Width is the width of the image. We want the precalculated and stored to avoid continually calculating this.
	Width uint16 `json:"width"`

	// Height is the height of the image. We want the precalculated and stored to avoid continually calculating this.
	Height uint16 `json:"height"`
}

// TextStyle defines styling and positioning for text.
type TextStyle struct {
	// X is the x-axis value representing the top-left of a text field.
	X uint16 `json:"x"`

	// Y is the y-axis value representing the top-left of a text field.
	Y uint16 `json:"y"`

	// Width is the width of the text field.
	Width uint16 `json:"width"`

	// Font is the font styling for the text.
	Font Font `json:"font"`

	// Stroke is the stroke/outline for the text.
	Stroke *Stroke `json:"stroke"`

	// Rotation is the rotation for the text.
	Rotation *Rotation `json:"rotation"`
}

// Font defines styling for text.
type Font struct {
	// Family defines the font family (e.g., "Helvetica").
	Family string `json:"family"`

	// Size defines the font size in point values.
	Size uint8 `json:"size"`

	// Color defines the color of the text in hexadecimal (e.g., "#FFFFFF").
	Color string `json:"color"`
}

// Stroke defines the stroke styling for text. It's the border around text.
type Stroke struct {
	// Size defines the size of the text stroke.
	Size uint8 `json:"size"`

	// Color defines the color of the text stroke in hexadecimal (e.g., "#FFFFFF").
	Color string `json:"color"`
}

// Rotation defines the rotation for text.
type Rotation struct {
	// Degrees defines the number of degrees to rotate the text.
	Degrees int16 `json:"degrees"`
}

// Config defines configuration for the Server.
type Config struct {
	httpapi.ServerConfig `mapstructure:"server"`

	// TemplateRepositoryType defines the GetterType for templates.
	TemplateRepositoryType RepositoryType `mapstructure:"template_repository_type"`

	// ImageGetterType defines the GetterType for template images.
	ImageGetterType image.GetterType `mapstructure:"image_getter_type"`

	// ImageUploaderType defines the ImageUploaderType.
	ImageUploaderType image.UploaderType `mapstructure:"image_uploader_type"`
}

type ConfigWrapper struct {
	Template Config `mapstructure:"template"`
}
