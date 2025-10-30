package imagekit

import (
	"embed"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/rotisserie/eris"
	"github.com/zq-xu/gotools/logx"
)

type Image interface {
	Image() *ebiten.Image
	GoImage() image.Image

	Width() int
	Height() int
}

// NewImageFromEmbed
func NewImageFromEmbed(embedFS *embed.FS, imgPath string) (Image, error) {
	f, err := embedFS.Open(imgPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, goImg, err := ebitenutil.NewImageFromReader(f)
	if err != nil {
		return nil, eris.Wrap(err, "new image from reader failed.")
	}

	logx.Logger.Debugf("embed image %s: width=%d, height=%d", imgPath, img.Bounds().Dx(), img.Bounds().Dy())

	return &basicImage{
		img:    img,
		goImg:  goImg,
		width:  img.Bounds().Dx(),
		height: img.Bounds().Dy(),
	}, nil
}

// NewImageListFromEmbed
func NewImageListFromEmbed(embedFS *embed.FS, imgPaths []string) ([]Image, error) {
	list := make([]Image, 0)

	for _, v := range imgPaths {
		img, err := NewImageFromEmbed(embedFS, v)
		if err != nil {
			return nil, eris.Wrap(err, "new image from embed failed")
		}
		list = append(list, img)
	}

	return list, nil
}

// NewImageFromFile
func NewImageFromFile(imgPath string) (Image, error) {
	img, goImg, err := ebitenutil.NewImageFromFile(imgPath)
	if err != nil {
		return nil, eris.Wrap(err, "new image from reader failed.")
	}

	logx.Logger.Debugf("embed image %s: width=%d, height=%d", imgPath, img.Bounds().Dx(), img.Bounds().Dy())

	return &basicImage{
		img:    img,
		goImg:  goImg,
		width:  img.Bounds().Dx(),
		height: img.Bounds().Dy(),
	}, nil
}
