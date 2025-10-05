package actor

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/internal/dungeon/config"
	"github.com/zq-xu/go-game/pkg/event/collision"
)

const stepLength = 1

type position struct {
	collisionObject collision.Object
}

func (a *actor) initPosition() {
	a.position = &position{
		collisionObject: collision.NewResolvRectObject("The knight", 100, 100, config.ActorWidth, config.ActorHeight),
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
