package uploader

// NewUploader constructs an Uploader based on the Type argument.
func NewUploader(t Type) Uploader {
	switch t {
	case Local:
		return NewLocalUploader()
	case Noop:
		return NewNoopUploader()
	}
	return NewLocalUploader()
}
