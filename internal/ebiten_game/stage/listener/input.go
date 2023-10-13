package listener

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/internal/ebiten_game/game"
)

type InputListener struct {
	ctx *game.Context

	counter int

	threshold int

	callback func() bool // return whether deal the input event
}

func NewInputListener(ctx *game.Context, callback func() bool) *InputListener {
	return &InputListener{
		ctx:       ctx,
		callback:  callback,
		threshold: int(ebiten.ActualTPS()) * ctx.Resource.Cfg.KeyInterval / 1000,
	}
}

func (g *InputListener) Update() {
	if g.callback == nil {
		return
	}

	if g.counter < g.threshold {
		g.counter++
		return
	}

	if g.callback() {
		g.counter = 0
	}
}
