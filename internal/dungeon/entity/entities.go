package entity

import (
	"github.com/rotisserie/eris"

	"github.com/zq-xu/go-game/internal/dungeon/core/actor"
	"github.com/zq-xu/go-game/internal/dungeon/core/tiledmap"
	"github.com/zq-xu/go-game/internal/dungeon/entity/monstors/skeleton"
	"github.com/zq-xu/go-game/internal/dungeon/entity/npc"
	"github.com/zq-xu/go-game/internal/dungeon/entity/player"
)

func NewEntities(tm tiledmap.TiledMap) ([]actor.Actor, error) {
	list := make([]actor.Actor, 0)
	p, err := player.NewPlayer()
	if err != nil {
		return nil, eris.Wrap(err, "failed to load player")
	}
	list = append(list, p)

	n, err := npc.LoadNPCs()
	if err != nil {
		return nil, eris.Wrap(err, "failed to load npcs")
	}
	list = append(list, n...)

	s, err := skeleton.NewSkeletonList(tm)
	if err != nil {
		return nil, eris.Wrap(err, "failed to load skeletons")
	}

	list = append(list, s...)
	return list, nil
}
