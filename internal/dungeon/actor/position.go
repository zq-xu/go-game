package actor

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/pkg/event/collision"
)

const stepLength = 1

type position struct {
	collisionObject collision.Object
}

func (a *actor) initPosition() {
	a.position = &position{
		collisionObject: collision.NewResolvRectObject("The knight", 100, 100, 20, 30),
	}
}

func (p *position) update(key ebiten.Key) {
	switch key {
	case ebiten.KeyUp:
		p.collisionObject.Move(0, -stepLength)
	case ebiten.KeyDown:
		p.collisionObject.Move(0, stepLength)
	case ebiten.KeyLeft:
		p.collisionObject.Move(-stepLength, 0)
	case ebiten.KeyRight:
		p.collisionObject.Move(stepLength, 0)
	}
}
