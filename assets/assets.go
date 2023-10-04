package assets

import (
	"embed"
)

//go:embed images
var EmbeddedImages embed.FS

//go:embed fonts
var EmbeddedFonts embed.FS
