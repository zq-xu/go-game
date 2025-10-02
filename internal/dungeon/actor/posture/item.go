package posture

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rotisserie/eris"

	"github.com/zq-xu/go-game/internal/dungeon/resources"
	"github.com/zq-xu/go-game/pkg/graphics/images/imagetable"
)

const chanegPostureInterval = 20

type postureItem interface {
	Update()
	Draw(screen *ebiten.Image, x, y float64)
}

type item struct {
	counter int // to change the image

	imageTable imagetable.RollImages
}

func newPostureItem(imgPath string) (postureItem, error) {
	s := &item{}

	img, err := resources.NewDungeonImage(imgPath)
	if err != nil {
		return nil, eris.Wrapf(err, "failed to load image file %s.", img)
	}

	s.imageTable = imagetable.NewRollImages(img)
	return s, nil
}

func (a *item) Update() {
	if a.counter < chanegPostureInterval {
		a.counter++
		return
	}

	a.counter = 0
	a.imageTable.Next()
}

func (a *item) Draw(screen *ebiten.Image, x, y float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)

	screen.DrawImage(a.imageTable.Image().Image(), op)
}
