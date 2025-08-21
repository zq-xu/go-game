package font

import (
	"bytes"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/rotisserie/eris"

	"github.com/zq-xu/go-game/assets"
)

// NewFontFromFile gets font from local file
func NewFontFromFile(path string) ([]byte, error) {
	return assets.EmbeddedFonts.ReadFile(path)
}

// LoadFont loads font
func LoadFont(size float64, ttf []byte) (text.Face, error) {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(ttf))
	if err != nil {
		return nil, eris.Wrap(err, "new font error")
	}

	return &text.GoTextFace{Source: s, Size: size}, nil
}
