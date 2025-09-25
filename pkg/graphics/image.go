package graphics

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/pkg/graphics/image"
)

var (
	NewImage = image.NewImage

	NewShooterImage = image.NewShooterImage
	NewDungeonImage = image.NewDungeonImage

	NewDungeonImageTable = image.NewDungeonImageTable

	GetImage = image.GetImage

	GetNineSliceImage       = image.GetNineSliceImage
	GetFixedNineSlice       = image.GetFixedNineSlice
	GetNineSliceSimpleImage = image.GetNineSliceSimpleImage
)

type Image interface {
	Image() *ebiten.Image
	Width() int
	Height() int
}
