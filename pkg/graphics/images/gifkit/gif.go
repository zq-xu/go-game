package gifkit

import (
	"embed"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rotisserie/eris"

	"github.com/zq-xu/go-game/pkg/graphics/images/imagekit"
)

type Gif interface {
	Draw(screen *ebiten.Image, x, y float64)

	IsDrawing() bool
	MoveToStart()

	SetUpdateInterval(i int)
}

type gifBase struct {
	updateInterval int
	counter        int

	index  int
	images []imagekit.Image
}

func newGifBase(imgs ...imagekit.Image) *gifBase {
	return &gifBase{
		index:  0,
		images: imgs,
	}
}

func newGifBaseFromEmbed(embedFS *embed.FS, imgPaths []string) (*gifBase, error) {
	list, err := imagekit.NewImageListFromEmbed(embedFS, imgPaths)
	if err != nil {
		return nil, eris.Wrap(err, "new image list from embed failed")
	}

	return newGifBase(list...), nil
}

func (g *gifBase) draw(screen *ebiten.Image, x, y float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	screen.DrawImage(g.images[g.index].Image(), op)
}

// func (g *gifBase) IsDrawFinished() bool {
// 	return g.index >= len(g.images)
// }

func (g *gifBase) MoveToStart() { g.index = 0 }

func (g *gifBase) next() {
	if g.counter < g.updateInterval {
		g.counter++
		return
	}

	g.index = (g.index + 1) % len(g.images)
	g.counter = 0
}

func (g *gifBase) SetUpdateInterval(i int) {
	g.updateInterval = i
}
