package img

import (
	_ "embed"
)

var (
	//go:embed window.png
	Window []byte

	//go:embed gophers.jpg
	Gophers []byte

	//go:embed bg.png
	Background []byte
)
