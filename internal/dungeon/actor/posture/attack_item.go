package posture

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rotisserie/eris"

	"github.com/zq-xu/go-game/internal/dungeon/resources"
	"github.com/zq-xu/go-game/pkg/graphics/images/imagetable"
)

type attackItem struct {
	direction  Direction
	rollImages imagetable.RollImages
}

func newAttackItem(direction Direction, imgPath string) (*attackItem, error) {
	s := &attackItem{direction: direction}

	img, err := resources.NewDungeonImage(imgPath)
	if err != nil {
		return nil, eris.Wrapf(err, "failed to load image file %s.", img)
	}

	s.rollImages = imagetable.NewRollImages(fmt.Sprintf("attack %d", direction), img)
	return s, nil
}

func (a *attackItem) Update() {
	a.rollImages.Next()
}

func (a *attackItem) Draw(screen *ebiten.Image, x, y float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	screen.DrawImage(a.rollImages.Image().Image(), op)
}

func (a *attackItem) Direction() Direction { return a.direction }

func (a *attackItem) MoveToStart() { a.rollImages.MoveToStart() }
