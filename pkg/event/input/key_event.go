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
	lastPress  time.Time
}

// return:
//  1. if pressed;
//  2. if exclusive
func (e *KeyEvent) listern(lastPress ebiten.Key) (bool, bool) {
	if !ebiten.IsKeyPressed(e.Key) {
		return false, false
	}

	// logs.Logger.Debugf("lastPress %s key %s", lastPress.String(), e.Key.String())
	if lastPress == e.Key && !e.Repeatable {
		return true, false
	}

	if time.Since(e.lastPress) >= e.Interval {
		e.Do()
		e.lastPress = time.Now()
	}

	return true, e.Exclusive
}
