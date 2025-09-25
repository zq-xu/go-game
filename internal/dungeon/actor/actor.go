package actor

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rotisserie/eris"

	"github.com/zq-xu/go-game/internal/dungeon/actor/position"
	"github.com/zq-xu/go-game/internal/dungeon/actor/posture"
	"github.com/zq-xu/go-game/internal/dungeon/config"
)

type Actor interface {
	Update()
	Draw(screen *ebiten.Image)
}

type actor struct {
	postures posture.Postures

	position position.Position
}

func NewActor() (Actor, error) {
	var err error
	a := &actor{}

	a.postures, err = posture.NewPostures()
	if err != nil {
		return nil, eris.Wrap(err, "failed to load actor steps")
	}

	a.position = position.NewPosition(config.MapWidth, config.MapHeight)
	return a, nil
}

func (a *actor) IdleDown() {}

func (a *actor) Update() {
	a.position.Update()
	a.postures.Update()
}

func (a *actor) Draw(screen *ebiten.Image) {
	a.postures.Draw(screen, a.position.X(), a.position.Y())
}
