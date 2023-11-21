package listener

import "github.com/zq-xu/go-game/internal/ebiten_game/event"

type inputListener struct {
	input *event.Input
}

func NewInputListener() *inputListener {
	return &inputListener{input: event.NewInput()}
}

func (il *inputListener) Update() error {
	il.input.Update()
	return nil
}
