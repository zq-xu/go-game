package components

import (
	"github.com/ebitenui/ebitenui/image"

	"github.com/zq-xu/go-game/pkg/graphics"
)

const backgroundColor = "131a22"

func NewBackgroundImage() *image.NineSlice {
	return image.NewNineSliceColor(graphics.HexToColor(backgroundColor))
}
