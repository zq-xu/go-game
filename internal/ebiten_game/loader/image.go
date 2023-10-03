package loader

import (
	"log"

	"github.com/zq-xu/2d-game/internal/ebiten_game/resources"
	"github.com/zq-xu/2d-game/pkg/graphics"
)

type ImageLoader struct {
	ufoImage    *graphics.Image
	bulletImage *graphics.Image
	shipImage   *graphics.Image
}

func NewImageLoader() *ImageLoader {
	return &ImageLoader{}
}

func (il *ImageLoader) GetUFOImage() *graphics.Image {
	return getImage(&il.ufoImage, resources.UFOPng)
}

func (il *ImageLoader) GetBulletImage() *graphics.Image {
	return getImage(&il.bulletImage, resources.BulletPng)
}

func (il *ImageLoader) GetShipImage() *graphics.Image {
	return getImage(&il.shipImage, resources.ShipPng)
}

func getImage(p **graphics.Image, bs []byte) *graphics.Image {
	if *p == nil {
		img, err := graphics.NewImage(bs)
		if err != nil {
			log.Fatal(err)
		}

		*p = img
	}

	return *p
}
