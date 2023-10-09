package runtime

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"

	"github.com/zq-xu/2d-game/internal/ebiten_game/entity"
	"github.com/zq-xu/2d-game/internal/ebiten_game/game"
	"github.com/zq-xu/2d-game/pkg/metric"
)

const UFOsName = "UFOs"

type UFOs struct {
	ctx *game.Context

	lastAddAt time.Time
	ufos      map[*entity.UFO]bool

	// TODO move to config
	MaxUFONum   int           // the max UFO count
	UFOInterval time.Duration // the interval between two UFOs, the unit is ms.
}

func NewUFOs(ctx *game.Context) *UFOs {
	return &UFOs{
		ctx:         ctx,
		ufos:        make(map[*entity.UFO]bool, 0),
		UFOInterval: 1000 * time.Millisecond,
		MaxUFONum:   20,
	}
}

func (us *UFOs) Update() error {
	if len(us.ufos) < us.MaxUFONum &&
		time.Since(us.lastAddAt) >= us.UFOInterval {
		u := entity.NewUFO(us.ctx)
		us.AddUFO(u)
	}

	for u := range us.ufos {
		inScreen := u.Update()
		if !inScreen {
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

func (us *UFOs) DrawMetrics(screen *ebiten.Image, cfg *metric.DrawConfig) {
	text.Draw(screen, fmt.Sprintf("UFOsCount: %d\n", len(us.ufos)), cfg.Face, cfg.X, cfg.Y, cfg.Color)
}

func (us *UFOs) RangeUFOs(fn func(k *entity.UFO, v bool)) {
	for k, v := range us.ufos {
		fn(k, v)
	}
}
