package mageguardian

import (
	"image"

	"github.com/rotisserie/eris"

	"github.com/zq-xu/go-game/assets/dungeon"
	"github.com/zq-xu/go-game/internal/dungeon/core/actor"
	"github.com/zq-xu/go-game/pkg/event/input"
)

type mageGuardian struct {
	actor.Actor
}

type MageGuardianConfig struct {
	Name string

	StartPoint *image.Point
}

// NewSkeleton
func NewMageGuardian(cfg *MageGuardianConfig) (actor.Actor, error) {
	a, err := actor.NewActor(&actor.Option{
		Name:       cfg.Name,
		StartPoint: cfg.StartPoint,
		Width:      35,
		Height:     60,
		IdleImagePaths: map[input.Direction]string{
			input.DefaultDirection: dungeon.GetDungeonImagePath("/characters/npcs/mage_guardian/idle.png"),
		},
	})
	if err != nil {
		return nil, eris.Wrap(err, "failed to new actor")
	}

	mg := &mageGuardian{Actor: a}
	return mg, nil
}

func (s *mageGuardian) Update() {
	s.Idle()
}
