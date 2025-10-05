package moving

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rotisserie/eris"

	"github.com/zq-xu/go-game/internal/dungeon/config"
	"github.com/zq-xu/go-game/internal/dungeon/resources"
	"github.com/zq-xu/go-game/pkg/graphics/images/imagetable"
)

const chanegPostureInterval = 20

type movingItem struct {
	counter    int // to change the image
	direction  config.Direction
	imageTable imagetable.RollImages
}

func newMovingItem(direction config.Direction, imgPath string) (*movingItem, error) {
	s := &movingItem{direction: direction}

	img, err := resources.NewDungeonImage(imgPath)
	if err != nil {
		return nil, eris.Wrapf(err, "failed to load image file %s.", img)
	}

	s.imageTable = imagetable.NewRollImages(fmt.Sprintf("moving %d", direction), img)
	return s, nil
}

func (a *movingItem) Update() {
	if a.counter < chanegPostureInterval {
		a.counter++
		return
	}

	a.counter = 0
	a.imageTable.Next()
}

func (a *movingItem) Draw(screen *ebiten.Image, x, y float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	screen.DrawImage(a.imageTable.Image().Image(), op)
}

func (a *movingItem) Direction() config.Direction { return a.direction }
