package actors

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"

	"github.com/zq-xu/go-game/internal/stages/gaming/gamerun/actors/entity"
	"github.com/zq-xu/go-game/pkg/metrics"
)

const UFOsName = "UFOs"

type UFOs struct {
	lastAddAt time.Time
	ufos      map[*entity.UFO]bool

	// TODO move to config
	MaxUFONum   int           // the max UFO count
	UFOInterval time.Duration // the interval between two UFOs, the unit is ms.
}

func NewUFOs() *UFOs {
	return &UFOs{

		ufos:        make(map[*entity.UFO]bool, 0),
		UFOInterval: 1000 * time.Millisecond,
		MaxUFONum:   20,
	}
}

/*
Update: add ufos per interval.
if the ufo is down of the screen, remove it from the storage.
*/
func (us *UFOs) Update() error {
	if len(us.ufos) < us.MaxUFONum &&
		time.Since(us.lastAddAt) >= us.UFOInterval {
		u := entity.NewUFO()
		us.AddUFO(u)
	}

	for u := range us.ufos {
		u.Update()

		if u.IsDownOfScreen() {
			us.RemoveUFO(u)
		}
	}

	return nil
}

func (us *UFOs) AddUFO(ufo *entity.UFO) {
	us.ufos[ufo] = true
	us.lastAddAt = time.Now()
}

func (us *UFOs) RemoveUFO(u *entity.UFO) {
	delete(us.ufos, u)
}

func (us *UFOs) Draw(screen *ebiten.Image) {
	for u := range us.ufos {
		u.Draw(screen)
	}
}

func (us *UFOs) DrawMetrics(screen *ebiten.Image, dc *metrics.DrawConfig) {
	op := &text.DrawOptions{}

	op.GeoM.Translate(float64(dc.X), float64(dc.Y))
	op.ColorScale.ScaleWithColor(dc.Color)
	text.Draw(screen, fmt.Sprintf("%s: UFOsCount: %d", UFOsName, len(us.ufos)), dc.Face, op)
}

func (us *UFOs) RangeUFOs(fn func(k *entity.UFO, v bool)) {
	for k, v := range us.ufos {
		fn(k, v)
	}
}
