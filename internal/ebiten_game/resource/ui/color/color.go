package color

import (
	"image/color"

	"github.com/zq-xu/2d-game/pkg/graphics"
)

const (
	textIdleRGBA     = "dff4ff"
	textDisabledRGBA = "5a7a91"
)

type ColorResource struct {
	TextIdleColor     color.Color
	TextDisabledColor color.Color
}

func NewColorResource() *ColorResource {
	return &ColorResource{
		TextIdleColor:     graphics.HexToColor(textIdleRGBA),
		TextDisabledColor: graphics.HexToColor(textDisabledRGBA),
	}
}
