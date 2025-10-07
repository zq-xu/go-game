package characters

import (
	"github.com/rotisserie/eris"
	"github.com/zq-xu/go-game/internal/dungeon/characters/monstors/skeleton"
	"github.com/zq-xu/go-game/internal/dungeon/characters/player"
	"github.com/zq-xu/go-game/internal/dungeon/core/actor"
)

type Character interface {
	actor.Actor

	Update()
}

func NewCharacters() ([]Character, error) {
	list := make([]Character, 0)
	p, err := player.NewPlayer()
	if err != nil {
		return nil, eris.Wrap(err, "failed to load player")
	}
	list = append(list, p)

	s, err := skeleton.NewSkeleton()
	if err != nil {
		return nil, eris.Wrap(err, "failed to load skeletons")
	}

	list = append(list, s)
	return list, nil
}
