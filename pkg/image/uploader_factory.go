package image

// NewUploader constructs an UploaderType based on the Type argument.
func NewUploader(t UploaderType) Uploader {
	switch t {
	case LocalUploader:
		return NewLocalUploader()
	case NoopUploader:
		return NewNoopUploader()
	}
	return NewLocalUploader()
}
