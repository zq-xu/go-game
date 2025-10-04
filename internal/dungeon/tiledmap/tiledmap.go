package tiledmap

import (
	"bytes"
	"fmt"

	"github.com/lafriks/go-tiled"
	"github.com/rotisserie/eris"

	"github.com/zq-xu/go-game/pkg/event/collision"
	"github.com/zq-xu/go-game/pkg/logs"
)

type TiledMap interface {
	Background

	Width() int
	Height() int

	TileWidth() int
	TileHeight() int

	PrintTiles()
	PrintObjects()

	CollisionObjects() []collision.Object
}

type tiledMap struct {
	tMap *tiled.Map

	Background

	Tiles
}

func NewTiledMap(bgImgPath, mapPath string) (TiledMap, error) {
	var err error
	var tm tiledMap

	tm.tMap, err = tiled.LoadFile(mapPath)
	if err != nil {
		return nil, eris.Wrap(err, "failed to load map")
	}

	tm.Background, err = NewBackground(bgImgPath)
	if err != nil {
		return nil, eris.Wrap(err, "failed to load map background")
	}

	tm.initTiles()
	return &tm, nil
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
