package tiledmap

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rotisserie/eris"

	"github.com/zq-xu/go-game/internal/dungeon/resources"
	"github.com/zq-xu/go-game/pkg/graphics"
)

type imgBackground struct {
	img *ebiten.Image
}

func newBackground(imgPath string) (*imgBackground, error) {
	var err error
	var b imgBackground

	basicImg, err := resources.NewDungeonImage(imgPath)
	if err != nil {
		return nil, eris.Wrapf(err, "failed to load background image from file %s", imgPath)
	}

	b.img = basicImg.Image()
	return &b, nil
}

func (b *imgBackground) Draw(screen *ebiten.Image, x, y float64) {
	graphics.DrawImage(screen, b.img, x, y)
}
