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
	return time.Millisecond * 200
}

func (a *actor) initInputListener() error {
	a.inputListener = input.NewInputListener()

	// priority
	a.inputListener.Listen(&input.KeyEvent{
		Key: ebiten.KeySpace,
		Do: func() {
			if a.postures.IsAttacking() {
				return
			}
			a.postures.UpdateKeyPress(ebiten.KeySpace)
		},
		Interval:   getAttackInterval(),
		Exclusive:  true,
		Repeatable: false,
	})

	for _, key := range moveKeyList {
		a.inputListener.Listen(&input.KeyEvent{
			Key: key,
			Do: func() {
				if a.postures.IsAttacking() {
					return
				}
				a.move(key)
			},
			Exclusive:  true,
			Repeatable: true,
		})
	}

	a.inputListener.Idle(func() { a.move(-1) })
	return nil
}
