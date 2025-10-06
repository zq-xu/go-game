package actor

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/internal/dungeon/config"
	"github.com/zq-xu/go-game/pkg/event/input"
)

func getAttackInterval() time.Duration {
	return time.Millisecond * 200
}

func (a *actor) initInputListener() error {
	a.inputListener = input.NewInputListener()

	a.inputListener.Blocks(a.attack.IsAttack)

	// listen attack, highest priority
	a.inputListener.Listen(&input.KeyEvent{
		Key: input.AttackKey,
		Do: func() {
			a.attack.Attack(a.direction)
			a.status = config.AttackActorStatus
		},
		Interval:   getAttackInterval(),
		Exclusive:  true,
		Repeatable: false,
	})

	// listen moving on direction
	input.RangeKeyDirections(func(key ebiten.Key, direction input.Direction) {
		a.inputListener.Listen(&input.KeyEvent{
			Key: key,
			Do: func() {
				a.running.Update(key)
				a.direction = direction
				a.status = config.RunningActorStatus
			},
			Exclusive:  true,
			Repeatable: true,
		})
	})

	// listen idle
	a.inputListener.Idle(func() {
		a.idle.Update(a.direction)
		a.status = config.IdleActorStatus
	})
	return nil
}
