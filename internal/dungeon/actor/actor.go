package actor

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rotisserie/eris"

	"github.com/zq-xu/go-game/internal/dungeon/actor/posture"
	"github.com/zq-xu/go-game/pkg/event/collision"
	"github.com/zq-xu/go-game/pkg/event/input"
)

type Actor interface {
	Update()
	Draw(screen *ebiten.Image)

	CollisionObject() collision.Object
}

type actor struct {
	postures      posture.Postures
	position      *position
	inputListener input.InputListener
}

func NewActor() (Actor, error) {
	var err error
	a := &actor{}

	a.postures, err = posture.NewPostures()
	if err != nil {
		return nil, eris.Wrap(err, "failed to load actor steps")
	}

	a.initPosition()
	a.initInputListener()
	return a, nil
}

func (a *actor) Update() {
	a.inputListener.Update()
}

func (a *actor) move(key ebiten.Key) {
	a.position.update(key)
	a.postures.UpdateKeyPress(key)
}

func (a *actor) Draw(screen *ebiten.Image) {
	x, y := a.position.collisionObject.LeftTop()
	a.postures.Draw(screen, x, y)
}

func (a *actor) CollisionObject() collision.Object { return a.position.collisionObject }
