package tiledmap

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rotisserie/eris"

	"github.com/zq-xu/go-game/assets/dungeon"
	"github.com/zq-xu/go-game/pkg/graphics/images/gifkit"
)

type gifs struct {
	gifs map[string]gifItem
}

type gifItem struct {
	gifkit.Gif
	x, y float64
}

func (tm *tiledMap) initGifs() error {
	tm.gifs = &gifs{gifs: make(map[string]gifItem, 0)}

	for name, v := range tm.cfg.Gifs {
		gg, err := gifkit.NewNeverStopGifFromEmbed(&dungeon.EmbeddedDungeon, v.ImgPaths)
		if err != nil {
			return eris.Wrapf(err, "new embed gif for %s failed.", name)
		}

		tm.gifs.gifs[name] = gifItem{Gif: gg, x: v.X, y: v.Y}
	}
	return nil
}

func (g *gifs) Draw(screen *ebiten.Image, x, y float64) {
	for _, v := range g.gifs {
		v.Draw(screen, v.x, v.y)
	}
}
