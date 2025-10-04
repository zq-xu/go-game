package actor

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rotisserie/eris"

	"github.com/zq-xu/go-game/internal/dungeon/actor/posture"
	"github.com/zq-xu/go-game/pkg/event/collision"
)

type Actor interface {
	Update()
	Draw(screen *ebiten.Image)

	CollisionObject() collision.Object
}

type actor struct {
	postures posture.Postures
	position *position
}

func NewActor() (Actor, error) {
	var err error
	a := &actor{}

	a.postures, err = posture.NewPostures()
	if err != nil {
		return nil, eris.Wrap(err, "failed to load actor steps")
	}

	a.position = newPosition()
	return a, nil
}

func (a *actor) Update() {
	a.position.update()
	a.postures.Update()
}

func (a *actor) Draw(screen *ebiten.Image) {
	a.postures.Draw(screen,
		a.position.collisionObject.X(),
		a.position.collisionObject.Y())
}

func (a *actor) CollisionObject() collision.Object { return a.position.collisionObject }
