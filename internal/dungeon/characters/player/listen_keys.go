package player

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/pkg/event/input"
)

func getAttackInterval() time.Duration {
	return time.Millisecond * 200
}

func (p *player) initInputListener() error {
	p.inputListener = input.NewInputListener()

	p.inputListener.Blocks(p.IsAttack)

	// listen attack, highest priority
	p.inputListener.Listen(&input.KeyEvent{
		Key:        input.AttackKey,
		Do:         p.Attack,
		Interval:   getAttackInterval(),
		Exclusive:  true,
		Repeatable: false,
	})

	// listen moving on direction
	input.RangeKeyDirections(func(key ebiten.Key, direction input.Direction) {
		p.inputListener.Listen(&input.KeyEvent{
			Key:        key,
			Do:         func() { p.Moving(key) },
			Exclusive:  true,
			Repeatable: true,
		})
	})

	// listen idle
	p.inputListener.Idle(p.Idle)
	return nil
}
