package actor

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/pkg/event/input"
)

var (
	moveKeyList = []ebiten.Key{
		ebiten.KeyUp,
		ebiten.KeyArrowDown,
		ebiten.KeyArrowLeft,
		ebiten.KeyArrowRight,
	}
)

func getAttackInterval() time.Duration {
	return time.Millisecond * 100
}

func (a *actor) initInputListener() error {
	a.inputListener = input.NewInputListener()

	// priority
	a.inputListener.Listen(&input.KeyEvent{
		Key: ebiten.KeySpace,
		Do: func() {
			a.postures.UpdateKeyPress(ebiten.KeySpace)
		},
		Interval:   getAttackInterval(),
		Exclusive:  false,
		Repeatable: true,
	})

	for _, key := range moveKeyList {
		a.inputListener.Listen(&input.KeyEvent{
			Key:        key,
			Do:         func() { a.move(key) },
			Exclusive:  false,
			Repeatable: true,
		})
	}

	a.inputListener.Idle(a.postures.UpdateIdle)
	return nil
}
