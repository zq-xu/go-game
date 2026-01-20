package player

import (
	"slices"

	"github.com/zq-xu/go-game/internal/dungeon/core/controls"
)

type approachingObjects struct {
	npcs []string
}

func (p *player) initApproachingObjects() {
	p.approaching = &approachingObjects{
		npcs: make([]string, 0),
	}
}

func (p *player) RefreshApproaching() {
	list := p.Approaching()

	for _, v := range list {
		controls.ApproachingNPC(v)
	}

	for _, v := range p.approaching.npcs {
		if slices.Contains(list, v) {
			continue
		}
		controls.MoveAwaryNPC(v)
	}

	p.approaching.npcs = list
}
