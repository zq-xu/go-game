package tiledmap

import (
	"github.com/lafriks/go-tiled"
	"github.com/rotisserie/eris"
)

type TiledMap interface {
	Background

	PrintTiles()
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

	tm.loadTiles()
	tm.PrintTiles()
	return &tm, nil
}
