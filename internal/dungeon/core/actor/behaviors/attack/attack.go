package attack

import (
	"fmt"

	"github.com/rotisserie/eris"

	"github.com/zq-xu/go-game/internal/dungeon/resources"
	"github.com/zq-xu/go-game/pkg/event/input"
	"github.com/zq-xu/go-game/pkg/graphics/images/gifkit"
	"github.com/zq-xu/go-game/pkg/graphics/images/imagekit"
	"github.com/zq-xu/go-game/pkg/graphics/images/imagetable"
)

type Attack interface {
	Attack(d input.Direction)
	Image() imagekit.Image
	IsAttack() bool
}

type attack struct {
	direction     input.Direction
	gifkit.Gif    // the current one
	attackItemSet map[input.Direction]gifkit.Gif
}

func NewAttack(name string, attackImagesPath map[input.Direction]string) (Attack, error) {
	ap := &attack{
		direction:     input.DefaultDirection,
		attackItemSet: make(map[input.Direction]gifkit.Gif),
	}

	for k, v := range attackImagesPath {
		img, err := resources.NewDungeonImage(v)
		if err != nil {
			return nil, eris.Wrapf(err, "failed to new dungeon image %s", v)
		}
		imgTable := imagetable.NewImageTable(fmt.Sprintf("%s-attack-%s", name, k.String()), img)
		imgTable.LogBoxes()
		g := gifkit.NewOnceGif(imgTable.Images()...)
		g.SetUpdateInterval(3)

		ap.attackItemSet[k] = g
	}

	ap.Gif = ap.attackItemSet[ap.direction]
	return ap, nil
}

func (ap *attack) Attack(d input.Direction) {
	if d.Unknown() {
		return
	}

	if d != ap.direction {
		ap.Gif = ap.attackItemSet[d]
		ap.direction = d
	}

	ap.Gif.MoveToStart()
}

func (ap *attack) IsAttack() bool { return ap.IsDrawing() }
