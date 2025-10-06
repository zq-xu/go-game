package running

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rotisserie/eris"

	"github.com/zq-xu/go-game/assets/dungeon"
	"github.com/zq-xu/go-game/internal/dungeon/config"
	"github.com/zq-xu/go-game/internal/dungeon/resources"
	"github.com/zq-xu/go-game/pkg/event/collision"
	"github.com/zq-xu/go-game/pkg/event/input"
	"github.com/zq-xu/go-game/pkg/graphics/images/gifkit"
	"github.com/zq-xu/go-game/pkg/graphics/images/imagekit"
	"github.com/zq-xu/go-game/pkg/graphics/images/imagetable"
)

var (
	runImagesPath = map[input.Direction]string{
		input.UpDirection:    dungeon.GetDungeonImagePath("/actor/run/run_up.png"),
		input.DownDirection:  dungeon.GetDungeonImagePath("/actor/run/run_down.png"),
		input.LeftDirection:  dungeon.GetDungeonImagePath("/actor/run/run_left.png"),
		input.RightDirection: dungeon.GetDungeonImagePath("/actor/run/run_right.png"),
	}
)

type Running interface {
	Update(key ebiten.Key)
	Image() imagekit.Image
}

type running struct {
	direction input.Direction

	gifkit.Gif
	runImages map[input.Direction]gifkit.Gif

	obj collision.Object
}

func NewRunning(obj collision.Object) (Running, error) {
	rp := &running{
		obj:       obj,
		direction: input.DefaultDirection,
		runImages: make(map[input.Direction]gifkit.Gif),
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
	rp.obj.MoveByKey(key, config.StepLength)

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
