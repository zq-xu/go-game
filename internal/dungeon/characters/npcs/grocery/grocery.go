package grocery

import (
	"image"

	"github.com/rotisserie/eris"

	"github.com/zq-xu/go-game/assets/dungeon"
	"github.com/zq-xu/go-game/internal/dungeon/core/actor"
	"github.com/zq-xu/go-game/pkg/event/input"
)

type grocery struct {
	actor.Actor
}

type GroceryConfig struct {
	Name string

	StartPoint *image.Point
}

// NewSkeleton
func NewGrocery(cfg *GroceryConfig) (actor.Actor, error) {
	a, err := actor.NewActor(&actor.Option{
		Name:       cfg.Name,
		StartPoint: cfg.StartPoint,
		Width:      35,
		Height:     60,
		IdleImagePaths: map[input.Direction]string{
			input.DefaultDirection: dungeon.GetDungeonImagePath("/characters/npcs/grocery/idle.png"),
		},
	})
	if err != nil {
		return nil, eris.Wrap(err, "failed to new actor")
	}

	mg := &grocery{Actor: a}
	return mg, nil
}

func (s *grocery) Update() {
	s.Idle()
}
