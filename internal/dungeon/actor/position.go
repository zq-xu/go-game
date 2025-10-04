package actor

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/pkg/event/collision"
)

const stepLength = 1

type position struct {
	// minX, minY float64
	// maxX, maxY float64

	collisionObject collision.Object
}

func newPosition() *position {
	return &position{
		// minX: 0,
		// minY: 0,
		// maxX: maxX,
		// maxY: maxY,

		collisionObject: collision.NewResolvObject(100, 100, 20, 30),
	}
}

func (p *position) update() {
	switch {
	case ebiten.IsKeyPressed(ebiten.KeyUp):
		// p.y = lo.Max([]float64{p.y - stepLength, p.minY})
		p.collisionObject.Move(0, -stepLength)
	case ebiten.IsKeyPressed(ebiten.KeyDown):
		// p.y = lo.Min([]float64{p.y + stepLength, p.maxY})
		p.collisionObject.Move(0, stepLength)
	case ebiten.IsKeyPressed(ebiten.KeyLeft):
		// p.x = lo.Max([]float64{p.x - stepLength, p.minX})
		p.collisionObject.Move(-stepLength, 0)
	case ebiten.IsKeyPressed(ebiten.KeyRight):
		// p.x = lo.Min([]float64{p.x + stepLength, p.maxX})
		p.collisionObject.Move(stepLength, 0)
	}
}
