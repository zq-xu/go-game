package tiledmap

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/rotisserie/eris"
)

type Background interface {
	DrawBackground(screen *ebiten.Image, x, y float64)
}

type imgBackground struct {
	img *ebiten.Image
}

func NewBackground(imgPath string) (Background, error) {
	var err error
	var b imgBackground

	b.img, _, err = ebitenutil.NewImageFromFile(imgPath)
	if err != nil {
		return nil, eris.Wrapf(err, "failed to load background image from file %s", imgPath)
	}

	return &b, nil
}

func (b *imgBackground) DrawBackground(screen *ebiten.Image, x, y float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	screen.DrawImage(b.img, op)
}
