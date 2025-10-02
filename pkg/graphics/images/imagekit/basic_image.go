package imagekit

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

// NewBasicImageFromGoImage
func NewBasicImageFromGoImage(golImage image.Image) *basicImage {
	ebitenImage := ebiten.NewImageFromImage(golImage)
	return &basicImage{
		img:    ebitenImage,
		goImg:  golImage,
		width:  ebitenImage.Bounds().Dx(),
		height: ebitenImage.Bounds().Dy(),
	}
}

func (bi *basicImage) Image() *ebiten.Image { return bi.img }
func (bi *basicImage) GoImage() image.Image { return bi.goImg }
func (bi *basicImage) Width() int           { return bi.width }
func (bi *basicImage) Height() int          { return bi.height }
