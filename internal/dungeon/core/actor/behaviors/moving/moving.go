package moving

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rotisserie/eris"

	"github.com/zq-xu/go-game/internal/dungeon/resources"
	"github.com/zq-xu/go-game/pkg/event/collision"
	"github.com/zq-xu/go-game/pkg/event/input"
	"github.com/zq-xu/go-game/pkg/graphics/images/gifkit"
	"github.com/zq-xu/go-game/pkg/graphics/images/imagekit"
	"github.com/zq-xu/go-game/pkg/graphics/images/imagetable"
)

type Moving interface {
	Move(key ebiten.Key)
	Image() imagekit.Image
}

type moving struct {
	direction input.Direction

	gifkit.Gif
	movingImages map[input.Direction]gifkit.Gif

	obj        collision.Object
	stepLength int
}

func NewMoving(name string, obj collision.Object, stepLength int, movingImagesPath map[input.Direction]string) (Moving, error) {
	rp := &moving{
		obj:          obj,
		stepLength:   stepLength,
		direction:    input.DefaultDirection,
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
		g.SetUpdateInterval(5)

		rp.movingImages[k] = g
	}

	rp.Gif = rp.movingImages[rp.direction]
	return rp, nil
}

func (rp *moving) Move(key ebiten.Key) {
	rp.obj.MoveByKey(key, float64(rp.stepLength))

	d := input.GetDirection(key)
	if d.Unknown() {
		return
	}

	if rp.direction != d {
		rp.Gif = rp.movingImages[d]
		rp.direction = d
	}
}

func (rp *moving) IsDrawing() bool { return false }
