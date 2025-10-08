package idle

import (
	"fmt"

	"github.com/rotisserie/eris"

	"github.com/zq-xu/go-game/internal/dungeon/resources"
	"github.com/zq-xu/go-game/pkg/event/input"
	"github.com/zq-xu/go-game/pkg/graphics/images/gifkit"
	"github.com/zq-xu/go-game/pkg/graphics/images/imagekit"
	"github.com/zq-xu/go-game/pkg/graphics/images/imagetable"
)

const idleInterval = 3

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

func NewIdle(name string, idleImagesPath map[input.Direction]string) (Idle, error) {
	mp := &idle{
		direction:  input.DefaultDirection,
		idleImages: make(map[input.Direction]gifkit.Gif),
	}

	for k, v := range idleImagesPath {
		img, err := resources.NewDungeonImage(v)
		if err != nil {
			return nil, eris.Wrapf(err, "failed to new dungeon image %s", v)
		}
		imgTable := imagetable.NewImageTable(fmt.Sprintf("%s-idle-%s", name, k.String()), img)
		imgTable.LogBoxes()
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
