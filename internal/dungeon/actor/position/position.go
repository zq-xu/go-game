package position

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/samber/lo"
)

const stepLength = 1

type Position interface {
	X() float64
	Y() float64
	Update()
}

type position struct {
	minX, minY, maxX, maxY float64
	x, y                   float64
}

func NewPosition(maxX, maxY float64) Position {
	return &position{
		minX: 0,
		minY: 0,
		maxX: maxX,
		maxY: maxY,

		x: 100,
		y: 100,
	}
}

func (p *position) Update() {
	switch {
	case ebiten.IsKeyPressed(ebiten.KeyUp):
		p.y = lo.Max([]float64{p.y - stepLength, p.minY})
	case ebiten.IsKeyPressed(ebiten.KeyDown):
		p.y = lo.Min([]float64{p.y + stepLength, p.maxY})
	case ebiten.IsKeyPressed(ebiten.KeyLeft):
		p.x = lo.Max([]float64{p.x - stepLength, p.minX})
	case ebiten.IsKeyPressed(ebiten.KeyRight):
		p.x = lo.Min([]float64{p.x + stepLength, p.maxX})
	}
}
func (p *position) X() float64 { return p.x }
func (p *position) Y() float64 { return p.y }
