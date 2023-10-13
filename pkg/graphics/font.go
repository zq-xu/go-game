package graphics

import (
	"github.com/zq-xu/go-game/assets"
)

func NewFontFromFile(path string) ([]byte, error) {
	return assets.EmbeddedFonts.ReadFile(path)
}
