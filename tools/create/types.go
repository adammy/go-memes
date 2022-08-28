package create

type ConfigType string

type CommandType string

type CommandFunc func() error

// FileConfig defines how to construct a file from a template.
type FileConfig struct {
	// Name is the name of the template. Needs to be the exact name of the template file.
	Name string

	// Tmpl is the path template file to use.
	Tmpl string

	// Dir is the output directory.
	Dir string

	// File is the output filename.
	File string

	// Data is data for the template.
	Data interface{}
}
