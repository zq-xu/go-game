package ui

import (
	"github.com/ebitenui/ebitenui/image"

	"github.com/zq-xu/go-game/pkg/graphics"
)

const backgroundColor = "131a22"

type BackgroundResource struct {
	Image *image.NineSlice
}

func NewBackgroundResource() *BackgroundResource {
	background := image.NewNineSliceColor(graphics.HexToColor(backgroundColor))

	return &BackgroundResource{
		Image: background,
	}
}
