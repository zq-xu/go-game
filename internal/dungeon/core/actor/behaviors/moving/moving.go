package moving

import (
	"fmt"

	"github.com/rotisserie/eris"

	"github.com/zq-xu/go-game/internal/dungeon/resources"
	"github.com/zq-xu/go-game/pkg/event/collision"
	"github.com/zq-xu/go-game/pkg/event/input"
	"github.com/zq-xu/go-game/pkg/graphics/images/gifkit"
	"github.com/zq-xu/go-game/pkg/graphics/images/imagekit"
	"github.com/zq-xu/go-game/pkg/graphics/images/imagetable"
)

type Moving interface {
	Move(d input.Direction)
	Image() imagekit.Image
}

type moving struct {
	direction input.Direction

	gifkit.Gif
	movingImages map[input.Direction]gifkit.Gif

	obj collision.Object
}

func NewMoving(name string, obj collision.Object, movingImagesPath map[input.Direction]string) (Moving, error) {
	mp := &moving{
		obj:          obj,
		movingImages: make(map[input.Direction]gifkit.Gif),
	}

	for k, v := range movingImagesPath {
		img, err := resources.NewDungeonImage(v)
		if err != nil {
			return nil, eris.Wrapf(err, "failed to new dungeon image %s", v)
		}
		imgTable := imagetable.NewImageTable(fmt.Sprintf("%s-moving-%s", name, k.String()), img)
		imgTable.LogBoxes()
		g := gifkit.NewNeverStopGif(imgTable.Images()...)
		g.SetUpdateInterval(10)

		mp.movingImages[k] = g
	}

	mp.direction = mp.getDefaultDirection()
	mp.Gif = mp.movingImages[mp.direction]
	return mp, nil
}

func (mp *moving) Move(d input.Direction) {
	mp.obj.MoveDirection(d)

	if d.Unknown() {
		return
	}

	mp.refreshDirection(d)
}

func (mp *moving) getDefaultDirection() input.Direction {
	_, ok := mp.movingImages[input.DefaultDirection]
	if ok {
		return input.DefaultDirection
	}

	for k := range mp.movingImages {
		return k
	}

	return input.DefaultDirection
}

func (mp *moving) refreshDirection(d input.Direction) {
	_, ok := mp.movingImages[d]
	if !ok {
		return
	}

	if mp.direction == d {
		return
	}

	mp.Gif = mp.movingImages[d]
	mp.direction = d
}

func (mp *moving) IsDrawing() bool { return false }
