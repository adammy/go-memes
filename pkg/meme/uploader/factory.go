package uploader

func NewUploader(t UploaderType) Uploader {
	switch t {
	case Local:
		return NewLocalUploader()
	case Noop:
		return NewNoopUploader()
	}
	return NewLocalUploader()
}
