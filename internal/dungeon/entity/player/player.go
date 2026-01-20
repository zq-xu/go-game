package player

import (
	"github.com/rotisserie/eris"
	"github.com/zq-xu/gotools/configx"
	"github.com/zq-xu/gotools/logx"

	"github.com/zq-xu/go-game/internal/dungeon/core/actor"
	"github.com/zq-xu/go-game/pkg/event/input"
)

type Player interface {
	actor.Actor

	Update()
}

type player struct {
	actor.Actor

	approaching *approachingObjects

	// listen to the key input
	inputListener input.InputListener
}

var PlayerCfg actor.Option

func init() {
	configx.RegisterByFile("player", &PlayerCfg, configx.DefaultSetupFunc)
}

// NewPlayer
func NewPlayer() (Player, error) {
	logx.Logger.Info("Loading Player")
	a, err := actor.NewActor(&PlayerCfg)
	if err != nil {
		return nil, eris.Wrap(err, "failed to new actor")
	}

	p := &player{Actor: a}
	p.initInputListener()
	p.initApproachingObjects()

	logx.Logger.Info("Loaded Player")
	return p, nil
}

func (p *player) Update() {
	p.inputListener.Update()
	p.RefreshApproaching()
}
