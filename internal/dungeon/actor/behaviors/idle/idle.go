package idle

import (
	"fmt"

	"github.com/rotisserie/eris"

	"github.com/zq-xu/go-game/assets/dungeon"
	"github.com/zq-xu/go-game/internal/dungeon/resources"
	"github.com/zq-xu/go-game/pkg/event/input"
	"github.com/zq-xu/go-game/pkg/graphics/images/gifkit"
	"github.com/zq-xu/go-game/pkg/graphics/images/imagekit"
	"github.com/zq-xu/go-game/pkg/graphics/images/imagetable"
)

const idleInterval = 3

var (
	idleImagesPath = map[input.Direction]string{
		input.UpDirection:    dungeon.GetDungeonImagePath("/actor/idle/idle_up.png"),
		input.DownDirection:  dungeon.GetDungeonImagePath("/actor/idle/idle_down.png"),
		input.LeftDirection:  dungeon.GetDungeonImagePath("/actor/idle/idle_left.png"),
		input.RightDirection: dungeon.GetDungeonImagePath("/actor/idle/idle_right.png"),
	}
)

type Idle interface {
	Update(d input.Direction)
	Image() imagekit.Image
}

type idle struct {
	idleCounter int // to change to idles

	direction input.Direction
	gifkit.Gif
	idleImages map[input.Direction]gifkit.Gif
}

func NewIdle() (Idle, error) {
	mp := &idle{
		direction:  input.DefaultDirection,
		idleImages: make(map[input.Direction]gifkit.Gif),
	}

	for k, v := range idleImagesPath {
		img, err := resources.NewDungeonImage(v)
		if err != nil {
			return nil, eris.Wrapf(err, "failed to new dungeon image %s", v)
		}
		imgTable := imagetable.NewImageTable(fmt.Sprintf("idle-%s", k.String()), img)

		g := gifkit.NewNeverStopGif(imgTable.Images()...)
		g.SetUpdateInterval(5)

		mp.idleImages[k] = g
	}

	mp.Gif = mp.idleImages[mp.direction]
	return mp, nil
}

func (ip *idle) Update(d input.Direction) {
	if d.Unknown() {
		return
	}

	if ip.idleCounter < idleInterval {
		ip.idleCounter++
		return
	}

	if ip.direction != d {
		ip.idleCounter = 0
		ip.Gif = ip.idleImages[d]
		ip.direction = d
	}
}
