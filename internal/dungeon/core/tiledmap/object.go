package tiledmap

import (
	"github.com/lafriks/go-tiled"

	"github.com/zq-xu/go-game/pkg/event/collision"
)

func getObjectX(obj *tiled.Object) float64 { return obj.X }

func getObjectY(obj *tiled.Object) float64 {
	if obj.GID == 0 {
		return obj.Y
	}

	return obj.Y - obj.Height
}

func (tm *tiledMap) CollisionObjects() []collision.Object {
	list := make([]collision.Object, 0)

	for _, v := range tm.tMap.ObjectGroups {
		for _, obj := range v.Objects {
			o := collision.NewResolvRectObject(
				obj.Name,
				getObjectX(obj), getObjectY(obj),
				obj.Width, obj.Height,
			)
			list = append(list, o)
		}
	}

	return list
}
