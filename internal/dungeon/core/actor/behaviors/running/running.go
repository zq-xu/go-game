package running

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

type Running interface {
	Update(key ebiten.Key)
	Image() imagekit.Image
}

type running struct {
	direction input.Direction

	gifkit.Gif
	runImages map[input.Direction]gifkit.Gif

	obj        collision.Object
	stepLength int
}

func NewRunning(obj collision.Object, stepLength int, runImagesPath map[input.Direction]string) (Running, error) {
	rp := &running{
		obj:        obj,
		stepLength: stepLength,
		direction:  input.DefaultDirection,
		runImages:  make(map[input.Direction]gifkit.Gif),
	}

	for k, v := range runImagesPath {
		img, err := resources.NewDungeonImage(v)
		if err != nil {
			return nil, eris.Wrapf(err, "failed to new dungeon image %s", v)
		}
		imgTable := imagetable.NewImageTable(fmt.Sprintf("idle-%s", k.String()), img)

		g := gifkit.NewNeverStopGif(imgTable.Images()...)
		g.SetUpdateInterval(5)

		rp.runImages[k] = g
	}

	rp.Gif = rp.runImages[rp.direction]
	return rp, nil
}

func (rp *running) Update(key ebiten.Key) {
	rp.obj.MoveByKey(key, float64(rp.stepLength))

	d := input.GetDirection(key)
	if d.Unknown() {
		return
	}

	if rp.direction != d {
		rp.Gif = rp.runImages[d]
		rp.direction = d
	}
}

func (rp *running) IsDrawing() bool { return false }
