package image

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type basicImage struct {
	img *ebiten.Image

	goImg image.Image

	width  int
	height int
}

func (bi *basicImage) Image() *ebiten.Image { return bi.img }
func (bi *basicImage) GoImage() image.Image { return bi.goImg }
func (bi *basicImage) Width() int           { return bi.width }
func (bi *basicImage) Height() int          { return bi.height }
