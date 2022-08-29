package img

import (
	"errors"
)

var (
	ErrNotFound           = errors.New("image not found")
	ErrImgSizeTooLarge    = errors.New("image exceeds size limit")
	ErrInvalidContentType = errors.New("invalid content type")
	ErrInvalidImgFormat   = errors.New("invalid image format")
)
