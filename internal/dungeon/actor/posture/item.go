package posture

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rotisserie/eris"

	"github.com/zq-xu/go-game/pkg/graphics"
)

const chanegPostureInterval = 20

type postureItem interface {
	Update()
	Draw(screen *ebiten.Image, x, y float64)
}

type item struct {
	counter   int // to change the image
	stepIndex int

	imageTable []image.Image
}

func newPostureItem(img string) (postureItem, error) {
	s := &item{}

	it, err := graphics.NewDungeonImageTable(img)
	if err != nil {
		return nil, eris.Wrapf(err, "failed to load image table %s", img)
	}

	s.imageTable = it.Images()
	return s, nil
}

func (a *item) Update() {
	if a.counter < chanegPostureInterval {
		a.counter++
		return
	}

	a.counter = 0
	a.stepIndex = (a.stepIndex + 1) % len(a.imageTable)
}

func (a *item) Draw(screen *ebiten.Image, x, y float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	i := a.imageTable[a.stepIndex]
	screen.DrawImage(i.(*ebiten.Image), op)
}
