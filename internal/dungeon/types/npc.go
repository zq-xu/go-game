package types

import "github.com/zq-xu/go-game/internal/dungeon/core/actor"

type NPC interface {
	actor.Actor

	MoveNearby()
	MoveAway()
}
