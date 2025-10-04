package tiledmap

import (
	"bytes"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
	"github.com/rotisserie/eris"

	"github.com/zq-xu/go-game/internal/dungeon/config"
	"github.com/zq-xu/go-game/pkg/event/collision"
	"github.com/zq-xu/go-game/pkg/logs"
)

type TiledMap interface {
	Width() int
	Height() int

	TileWidth() int
	TileHeight() int

	PrintTiles()
	PrintObjects()

	CollisionObjects() []collision.Object

	Draw(screen *ebiten.Image, x, y float64)
}

type tiledMap struct {
	cfg *config.TiledMapConfig

	tMap *tiled.Map

	background *imgBackground
	gifs       *gifs

	Tiles
}

func NewTiledMap(cfg *config.TiledMapConfig) (TiledMap, error) {
	var err error
	tm := &tiledMap{cfg: cfg}

	tm.tMap, err = tiled.LoadFile(cfg.MapPath)
	if err != nil {
		return nil, eris.Wrap(err, "failed to load map")
	}

	tm.background, err = newBackground(cfg.BgImgPath)
	if err != nil {
		return nil, eris.Wrap(err, "failed to load map background")
	}

	err = tm.initGifs()
	if err != nil {
		return nil, eris.Wrap(err, "failed to init gifs")
	}

	tm.initTiles()
	return tm, nil
}

func (tm *tiledMap) Width() int      { return tm.tMap.Width * tm.tMap.TileWidth }
func (tm *tiledMap) Height() int     { return tm.tMap.Height * tm.tMap.TileHeight }
func (tm *tiledMap) TileWidth() int  { return tm.tMap.TileWidth }
func (tm *tiledMap) TileHeight() int { return tm.tMap.TileHeight }

func (tm *tiledMap) PrintObjects() {
	for _, v := range tm.tMap.ObjectGroups {
		for _, obj := range v.Objects {
			logs.Logger.Debugf("obj %s:\n%+v", obj.Name, obj)
		}
	}
}

func (tm *tiledMap) PrintTiles() {
	var buf bytes.Buffer

	for _, iv := range tm.Tiles {
		for _, jv := range iv {
			if jv.ObjectLength() > 0 {
				fmt.Fprintf(&buf, " %v", jv.ObjectLength())
			} else {
				buf.WriteString("  ")
			}
		}
		buf.WriteByte('\n')
	}

	fmt.Printf("tiledmap:\n%s\n", buf.String())
}

func (tm *tiledMap) Draw(screen *ebiten.Image, x, y float64) {
	tm.background.Draw(screen, 0, 0)
	tm.gifs.Draw(screen, 0, 0)
}
