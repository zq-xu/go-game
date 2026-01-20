package skeleton

import (
	"image"

	"github.com/rotisserie/eris"

	"github.com/zq-xu/go-game/internal/dungeon/entity/core/actor"
	"github.com/zq-xu/go-game/internal/dungeon/entity/core/controls"
	"github.com/zq-xu/go-game/internal/dungeon/entity/monstors/base"
)

type skeleton struct {
	actor.Actor

	m controls.Mover
}

type SkeletonConfig struct {
	Name        string
	StartPoint  *image.Point
	ActiveRange *image.Rectangle
}

var skeletonKey = "skeleton"

// NewSkeleton
func NewSkeleton(cfg *SkeletonConfig) (actor.Actor, error) {
	actorOption := base.MonsterConfigSet[skeletonKey].Option
	actorOption.Name = cfg.Name
	actorOption.StartPoint = cfg.StartPoint

	a, err := actor.NewActor(&actorOption)
	if err != nil {
		return nil, eris.Wrap(err, "failed to new actor")
	}

	s := &skeleton{Actor: a}
	s.m = controls.NewRandomMover(s, cfg.ActiveRange)
	return s, nil
}

func (s *skeleton) Update() {
	s.m.Move()
}
