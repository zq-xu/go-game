package gifkit

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zq-xu/go-game/pkg/graphics/images/imagekit"
)

type Gif interface {
	Draw(screen *ebiten.Image, x, y float64)
}

type gif struct {
	index  int
	images []imagekit.Image
}

func NewGIF(imgs ...imagekit.Image) Gif {
	it := &gif{
		index:  0,
		images: imgs,
	}

	return it
}

func (g *gif) Draw(screen *ebiten.Image, x, y float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	screen.DrawImage(g.images[g.index].Image(), op)
	g.index = (g.index + 1) % len(g.images)
}
