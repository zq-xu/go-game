package input

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type KeyEvent struct {
	Key        ebiten.Key
	Do         func()
	Exclusive  bool // If the key is pressed, ignore others
	Repeatable bool // Whether the key can be held down for continuous triggering.
	Interval   time.Duration

	lastPress time.Time
	pressed   bool
}

// return:
// 1. if pressed
// 2. if exclusive
func (e *KeyEvent) listern() (bool, bool) {
	if !ebiten.IsKeyPressed(e.Key) {
		e.pressed = false
		return e.pressed, false
	}

	// logs.Logger.Debugf("lastPress %s key %s", lastPress.String(), e.Key.String())
	if e.pressed && !e.Repeatable {
		return e.pressed, false
	}

	if time.Since(e.lastPress) >= e.Interval {
		e.Do()
		e.lastPress = time.Now()
	}

	e.pressed = true
	return e.pressed, e.Exclusive
}
