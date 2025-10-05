package gifkit

import (
	"embed"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/pkg/graphics/images/imagekit"
)

type neverStopGif struct {
	*gifBase
}

func NewNeverStopGif(imgs ...imagekit.Image) Gif {
	return &neverStopGif{
		gifBase: newGifBase(imgs...),
	}
}

func NewNeverStopGifFromEmbed(embedFS *embed.FS, imgPaths []string) (Gif, error) {
	g, err := newGifBaseFromEmbed(embedFS, imgPaths)
	if err != nil {
		return nil, err
	}

	return &neverStopGif{
		gifBase: g,
	}, nil
}

func (g *neverStopGif) Draw(screen *ebiten.Image, x, y float64) {
	g.gifBase.draw(screen, x, y)
	g.gifBase.next()
}

func (g *neverStopGif) IsDrawing() bool { return true }

func (g *neverStopGif) MoveToStart() { g.gifBase.MoveToStart() }
