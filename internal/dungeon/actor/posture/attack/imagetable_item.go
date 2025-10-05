package attack

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rotisserie/eris"

	"github.com/zq-xu/go-game/internal/dungeon/config"
	"github.com/zq-xu/go-game/internal/dungeon/resources"
	"github.com/zq-xu/go-game/pkg/graphics/images/imagetable"
)

type imageTableItem struct {
	direction  config.Direction
	rollImages imagetable.RollImages
}

func newImageTableItem(direction config.Direction, imgPath string) (*imageTableItem, error) {
	s := &imageTableItem{direction: direction}

	img, err := resources.NewDungeonImage(imgPath)
	if err != nil {
		return nil, eris.Wrapf(err, "failed to load image file %s.", img)
	}

	s.rollImages = imagetable.NewRollImages(fmt.Sprintf("attack %d", direction), img)
	return s, nil
}

func (a *imageTableItem) Draw(screen *ebiten.Image, x, y float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	screen.DrawImage(a.rollImages.Image().Image(), op)
	a.rollImages.Next()
}

func (a *imageTableItem) Direction() config.Direction { return a.direction }

func (a *imageTableItem) MoveToStart() { a.rollImages.MoveToStart() }

func (a *imageTableItem) isInProgress() bool { return true }

func (a *imageTableItem) IsInProgress() bool { return false }
