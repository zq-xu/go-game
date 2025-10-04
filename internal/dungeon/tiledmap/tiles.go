package tiledmap

import (
	"github.com/lafriks/go-tiled"
	"github.com/samber/lo"
)

type Tiles [][]Tile

type Tile interface {
	AddObject(obj *tiled.Object)
	ObjectLength() int
}

type tile struct {
	objects []*tiled.Object
}

func newTile() Tile {
	t := &tile{
		objects: make([]*tiled.Object, 0),
	}

	return t
}

func (t *tile) AddObject(obj *tiled.Object) {
	t.objects = append(t.objects, obj)
}

func (t *tile) ObjectLength() int { return len(t.objects) }

func (tm *tiledMap) initTiles() {
	tm.Tiles = make(Tiles, tm.tMap.Height)
	for i := 0; i < tm.tMap.Height; i++ {
		tm.Tiles[i] = make([]Tile, tm.tMap.Width)

		for j := 0; j < tm.tMap.Width; j++ {
			tm.Tiles[i][j] = newTile()
		}
	}

	// add objects to tiles
	for _, v := range tm.tMap.ObjectGroups {
		for _, obj := range v.Objects {
			tm.addObjectToTiles(obj)
		}
	}
}

func (tm *tiledMap) addObjectToTiles(obj *tiled.Object) {
	startX := int(obj.X) / tm.tMap.TileWidth
	endX := lo.Min([]int{int(obj.X+obj.Width) / tm.tMap.TileWidth, tm.tMap.Width - 1})

	startY := int(getObjectY(obj)) / tm.tMap.TileHeight
	endY := lo.Min([]int{int(getObjectY(obj)+obj.Height) / tm.tMap.TileHeight, tm.tMap.Height - 1})

	for i := startY; i <= endY; i++ {
		for j := startX; j <= endX; j++ {
			tm.Tiles[i][j].AddObject(obj)
		}
	}
}
