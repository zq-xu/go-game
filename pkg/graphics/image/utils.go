package image

import (
	"bytes"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/zq-xu/go-game/assets"
)

// NewDungeonImage
func NewDungeonImage(path string) (*basicImage, error) {
	f, err := assets.EmbeddedDungeon.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, goImg, err := ebitenutil.NewImageFromReader(f)
	if err != nil {
		return nil, err
	}

	return &basicImage{
		img:    img,
		goImg:  goImg,
		width:  img.Bounds().Dx(),
		height: img.Bounds().Dy(),
	}, nil
}

// NewShooterImage
func NewShooterImage(path string) (*basicImage, error) {
	f, err := assets.EmbeddedShooter.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, goImg, err := ebitenutil.NewImageFromReader(f)
	if err != nil {
		return nil, err
	}

	return &basicImage{
		img:    img,
		goImg:  goImg,
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
