package skeleton

import (
	"image"

	"github.com/rotisserie/eris"

	"github.com/zq-xu/go-game/assets/dungeon"
	"github.com/zq-xu/go-game/internal/dungeon/core/actor"
	"github.com/zq-xu/go-game/internal/dungeon/core/controls"
	"github.com/zq-xu/go-game/pkg/event/input"
)

type skeleton struct {
	actor.Actor

	direction input.Direction

	m controls.Mover
}

type SkeletonConfig struct {
	Name string

	StartPoint  *image.Point
	ActiveRange *image.Rectangle

	StepLength float64
}

// NewSkeleton
func NewSkeleton(cfg *SkeletonConfig) (actor.Actor, error) {
	a, err := actor.NewActor(&actor.Option{
		Name:       cfg.Name,
		StartPoint: cfg.StartPoint,
		StepLength: cfg.StepLength,
		Width:      36,
		Height:     33,
		AttackImagePaths: map[input.Direction]string{
			input.LeftDirection:  dungeon.GetDungeonImagePath("/characters/skeleton/attack/attack_left.png"),
			input.RightDirection: dungeon.GetDungeonImagePath("/characters/skeleton/attack/attack_right.png"),
		},
		IdleImagePaths: map[input.Direction]string{
			input.LeftDirection:  dungeon.GetDungeonImagePath("/characters/skeleton/idle/idle_left.png"),
			input.RightDirection: dungeon.GetDungeonImagePath("/characters/skeleton/idle/idle_right.png"),
		},
		MovingImagePaths: map[input.Direction]string{
			input.LeftDirection:  dungeon.GetDungeonImagePath("/characters/skeleton/moving/moving_left.png"),
			input.RightDirection: dungeon.GetDungeonImagePath("/characters/skeleton/moving/moving_right.png"),
		},
	})
	if err != nil {
		return nil, eris.Wrap(err, "failed to new actor")
	}

	s := &skeleton{
		Actor:     a,
		direction: input.RightDirection,
	}

	s.m = controls.NewRandomMover(s, cfg.ActiveRange)
	return s, nil
}

func (s *skeleton) Update() {
	s.m.Move()
}
