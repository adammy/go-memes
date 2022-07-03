package meme

// Meme defines a meme the user wants to make.
type Meme struct {
	ImgPath string
	Width   uint16
	Height  uint16
	Text    []Text
}

// Text defines text and it's associated styling and positioning, for a Meme.
type Text struct {
	X      uint16
	Y      uint16
	Width  uint16
	Text   string
	Font   Font
	Stroke Stroke
}

// Font defines styling for Text.
type Font struct {
	Family string
	Size   uint8
	Color  string
}

type Stroke struct {
	Enabled bool
	Size    uint8
	Color   string
}
