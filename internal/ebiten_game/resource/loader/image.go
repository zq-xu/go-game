package loader

import (
	"github.com/zq-xu/2d-game/internal/ebiten_game/resource/loader/image"
)

const (
	BulletImgPath = "images/bullet.png"
	ShipImgPath   = "images/ship.png"
	UFOImgPath    = "images/ufo.png"
)

type ImageLoader struct {
	ImgLoader *image.ImgLoader

	NineSliceImgLoader *image.NineSliceImgLoader
}

func NewImageLoader() *ImageLoader {
	return &ImageLoader{
		ImgLoader:          image.NewImgLoader(),
		NineSliceImgLoader: image.NewNineSliceImgLoader(),
	}
}
