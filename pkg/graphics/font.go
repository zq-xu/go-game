package graphics

import (
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/zq-xu/go-game/pkg/graphics/font"
)

var defaultFont Font

var (
	NewNotoSansFont = font.NewNotoSansFont
)

type Font interface {
	Face() *text.Face
	BoldFace() *text.Face
}

// GetFont
func GetFont() Font {
	if defaultFont == nil {
		defaultFont = NewNotoSansFont(12)
	}
	return defaultFont
}
