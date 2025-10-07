package player

import (
	"github.com/rotisserie/eris"

	"github.com/zq-xu/go-game/assets/dungeon"
	"github.com/zq-xu/go-game/internal/dungeon/config"
	"github.com/zq-xu/go-game/internal/dungeon/core/actor"
	"github.com/zq-xu/go-game/pkg/event/input"
)

type Player interface {
	actor.Actor

	Update()
}

type player struct {
	actor.Actor

	// listen to the key input
	inputListener input.InputListener
}

// NewPlayer
func NewPlayer() (Player, error) {
	a, err := actor.NewActor(&actor.Option{
		Name:   "Boy Knight",
		StartX: 100,
		StartY: 100,
		Width:  config.ActorWidth,
		Height: config.ActorHeight,
		AttackImagePaths: map[input.Direction]string{
			input.UpDirection:    dungeon.GetDungeonImagePath("/characters/player/attack/attack1_up.png"),
			input.DownDirection:  dungeon.GetDungeonImagePath("/characters/player/attack/attack1_down.png"),
			input.LeftDirection:  dungeon.GetDungeonImagePath("/characters/player/attack/attack1_left.png"),
			input.RightDirection: dungeon.GetDungeonImagePath("/characters/player/attack/attack1_right.png"),
		},
		IdleImagePaths: map[input.Direction]string{
			input.UpDirection:    dungeon.GetDungeonImagePath("/characters/player/idle/idle_up.png"),
			input.DownDirection:  dungeon.GetDungeonImagePath("/characters/player/idle/idle_down.png"),
			input.LeftDirection:  dungeon.GetDungeonImagePath("/characters/player/idle/idle_left.png"),
			input.RightDirection: dungeon.GetDungeonImagePath("/characters/player/idle/idle_right.png"),
		},
		MovingImagePaths: map[input.Direction]string{
			input.UpDirection:    dungeon.GetDungeonImagePath("/characters/player/run/run_up.png"),
			input.DownDirection:  dungeon.GetDungeonImagePath("/characters/player/run/run_down.png"),
			input.LeftDirection:  dungeon.GetDungeonImagePath("/characters/player/run/run_left.png"),
			input.RightDirection: dungeon.GetDungeonImagePath("/characters/player/run/run_right.png"),
		},
	})
	if err != nil {
		return nil, eris.Wrap(err, "failed to new actor")
	}

	p := &player{Actor: a}
	p.initInputListener()

	return p, nil
}

func (p *player) Update() {
	p.inputListener.Update()
}
