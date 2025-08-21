package graphics

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/pkg/graphics/image"
)

var (
	NewImage         = image.NewImage
	NewImageFromFile = image.NewImageFromFile

	GetImage = image.GetImage

	GetNineSliceImage       = image.GetNineSliceImage
	GetNineSliceSimpleImage = image.GetNineSliceSimpleImage
)

type Image interface {
	Image() *ebiten.Image
	Width() int
	Height() int
}
