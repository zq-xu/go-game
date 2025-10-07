package skeleton

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zq-xu/go-game/internal/dungeon/core/actor"
)

type mov struct {
	delta int
	minX  float64
	maxX  float64

	a actor.Actor
}

func NewMov(a actor.Actor, minX, maxX float64) *mov {
	return &mov{
		delta: 1,
		minX:  minX,
		maxX:  maxX,
		a:     a,
	}
}

func (m *mov) RandomMove() {
	x, _ := m.a.LeftTop()
	if x >= m.maxX {
		m.delta = -1
	} else if x < m.minX {
		m.delta = 1
	}

	switch m.delta {
	case 1:
		m.a.Moving(ebiten.KeyRight)
	case -1:
		m.a.Moving(ebiten.KeyLeft)
	default:
		m.a.Idle()
	}
}
