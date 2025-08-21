package event

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/internal/config"
)

type InputListener interface {
	Update()
	Reload()
}

type inputListener struct {
	counter int

	threshold int

	callback func() bool // return whether deal the input event
}

func NewInputListener(callback func() bool) InputListener {
	return &inputListener{
		callback:  callback,
		threshold: int(ebiten.ActualTPS()) * config.Cfg.KeyInterval / 1000,
	}
}

func (g *inputListener) Update() {
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

func (g *inputListener) Reload() { g.counter = 0 }
