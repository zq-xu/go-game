package imagekit

import (
	"embed"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/rotisserie/eris"
)

type Image interface {
	Image() *ebiten.Image
	GoImage() image.Image

	Width() int
	Height() int
}

// NewImageFromEmbed
func NewImageFromEmbed(embedFS *embed.FS, path string) (*basicImage, error) {
	f, err := embedFS.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, goImg, err := ebitenutil.NewImageFromReader(f)
	if err != nil {
		return nil, eris.Wrap(err, "new image from reader failed.")
	}

	return &basicImage{
		img:    img,
		goImg:  goImg,
		width:  img.Bounds().Dx(),
		height: img.Bounds().Dy(),
	}, nil
}
