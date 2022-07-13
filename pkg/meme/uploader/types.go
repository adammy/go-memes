package uploader

type UploaderType string

const (
	Local UploaderType = "Local"
	Noop  UploaderType = "Noop"
)
