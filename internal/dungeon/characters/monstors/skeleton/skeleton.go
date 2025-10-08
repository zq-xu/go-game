package skeleton

import (
	"github.com/rotisserie/eris"

	"github.com/zq-xu/go-game/assets/dungeon"
	"github.com/zq-xu/go-game/internal/dungeon/core/actor"
	"github.com/zq-xu/go-game/pkg/event/input"
)

type Skeleton interface {
	actor.Actor

	Update()
}

type skeleton struct {
	actor.Actor

	direction  input.Direction
	stepLength int

	m *mov
	// inputListener input.InputListener
}

// NewSkeleton
func NewSkeleton() (Skeleton, error) {
	a, err := actor.NewActor(&actor.Option{
		Name:   "Skeleton",
		StartX: 160,
		StartY: 100,
		Width:  36,
		Height: 33,
		AttackImagePaths: map[input.Direction]string{
			input.UpDirection:    dungeon.GetDungeonImagePath("/characters/skeleton/attack/attack_up.png"),
			input.DownDirection:  dungeon.GetDungeonImagePath("/characters/skeleton/attack/attack_down.png"),
			input.LeftDirection:  dungeon.GetDungeonImagePath("/characters/skeleton/attack/attack_left.png"),
			input.RightDirection: dungeon.GetDungeonImagePath("/characters/skeleton/attack/attack_right.png"),
		},
		IdleImagePaths: map[input.Direction]string{
			input.UpDirection:    dungeon.GetDungeonImagePath("/characters/skeleton/idle/idle_up.png"),
			input.DownDirection:  dungeon.GetDungeonImagePath("/characters/skeleton/idle/idle_down.png"),
			input.LeftDirection:  dungeon.GetDungeonImagePath("/characters/skeleton/idle/idle_left.png"),
			input.RightDirection: dungeon.GetDungeonImagePath("/characters/skeleton/idle/idle_right.png"),
		},
		MovingImagePaths: map[input.Direction]string{
			input.UpDirection:    dungeon.GetDungeonImagePath("/characters/skeleton/moving/moving_up.png"),
			input.DownDirection:  dungeon.GetDungeonImagePath("/characters/skeleton/moving/moving_down.png"),
			input.LeftDirection:  dungeon.GetDungeonImagePath("/characters/skeleton/moving/moving_left.png"),
			input.RightDirection: dungeon.GetDungeonImagePath("/characters/skeleton/moving/moving_right.png"),
		},
	})
	if err != nil {
		return nil, eris.Wrap(err, "failed to new actor")
	}

	s := &skeleton{
		Actor:      a,
		direction:  input.RightDirection,
		stepLength: 1,
	}
	s.m = NewMov(s, 100, 200)
	// s.initInputListener()
	return s, nil
}

func (s *skeleton) Update() {
	s.m.RandomMove()
	// s.Idle()
	// s.inputListener.Update()
}
