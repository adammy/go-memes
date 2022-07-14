package uploader

const (
	// Local denotes to upload items to the os filesystem.
	Local Type = "Local"

	// Noop denotes to not upload items and fake success.
	Noop Type = "Noop"
)
