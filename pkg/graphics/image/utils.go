package image

import (
	"bytes"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/zq-xu/go-game/assets"
)

// NewImageFromFile
func NewImageFromFile(path string) (*basicImage, error) {
	f, err := assets.EmbeddedImages.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := ebitenutil.NewImageFromReader(f)
	if err != nil {
		return nil, err
	}

	return &basicImage{
		img:    img,
		width:  img.Bounds().Dx(),
		height: img.Bounds().Dy(),
	}, nil
}

// NewImage
func NewImage(imgByte []byte) (*basicImage, error) {
	img, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(imgByte))
	if err != nil {
		return nil, err
	}

	return &basicImage{
		img:    img,
		width:  img.Bounds().Dx(),
		height: img.Bounds().Dy(),
	}, nil
}
