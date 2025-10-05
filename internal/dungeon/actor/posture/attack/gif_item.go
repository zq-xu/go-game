package attack

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rotisserie/eris"

	"github.com/zq-xu/go-game/internal/dungeon/config"
	"github.com/zq-xu/go-game/internal/dungeon/resources"
	"github.com/zq-xu/go-game/pkg/graphics/images/gifkit"
	"github.com/zq-xu/go-game/pkg/graphics/images/imagetable"
)

type gifItem struct {
	direction config.Direction
	gifkit.Gif
}

func newGifItem(direction config.Direction, imgPath string) (*gifItem, error) {
	s := &gifItem{direction: direction}

	img, err := resources.NewDungeonImage(imgPath)
	if err != nil {
		return nil, eris.Wrapf(err, "failed to load image file %s.", img)
	}

	imgTable := imagetable.NewDungeonImageTable(fmt.Sprintf("attack %d", direction), img)
	s.Gif = gifkit.NewOnceGif(imgTable.Images()...)
	s.Gif.SetUpdateInterval(3)
	return s, nil
}

func (g *gifItem) Draw(screen *ebiten.Image, x, y float64) {
	if g.direction == config.LeftDirection || g.direction == config.UpDirection {
		x, y = getDrawBeginningOnRightBottom(g.Gif.Image(), x, y)
	}

	g.Gif.Draw(screen, x, y)
}

func (g *gifItem) Direction() config.Direction { return g.direction }

func (g *gifItem) isInProgress() bool { return g.IsDrawing() }
