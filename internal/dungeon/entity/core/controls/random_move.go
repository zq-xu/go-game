package controls

import (
	"image"
	"math/rand"

	"github.com/samber/lo"

	"github.com/zq-xu/go-game/pkg/event/collision"
	"github.com/zq-xu/go-game/pkg/event/input"
)

const changeInterval = 300

type Mover interface {
	Move()
}

type randomMover struct {
	mr collision.Object

	activeRange *image.Rectangle

	deltaX, deltaY     int
	counterX, counterY int
}

func NewRandomMover(mr collision.Object, activeRange *image.Rectangle) Mover {
	return &randomMover{
		mr: mr,

		activeRange: activeRange,

		deltaX: 1,
		deltaY: 1,
	}
}

func (m *randomMover) Move() {
	x, y := m.mr.LeftTop()

	m.updateDeltaX(int(x))
	m.updateDeltaY(int(y))

	m.performRandomMove()
}

func (m *randomMover) performRandomMove() {
	[]func(){
		m.moveHorizontally,
		m.moveVertically,
	}[rand.Intn(2)]()
}

func (m *randomMover) updateDeltaX(x int) {
	switch {
	case x >= m.activeRange.Max.X:
		m.setDeltaX(-1)
	case x <= m.activeRange.Min.X:
		m.setDeltaX(1)
	default:
		m.incrementAndMaybeChangeX()
	}
}

func (m *randomMover) setDeltaX(delta int) {
	m.deltaX = delta
	m.counterX = 0
}

func (m *randomMover) incrementAndMaybeChangeX() {
	m.counterX++
	if m.counterX >= changeInterval {
		m.deltaX = lo.Ternary(rand.Intn(2) == 0, -1, 1)
		m.counterX = 0
	}
}

func (m *randomMover) updateDeltaY(y int) {
	switch {
	case y >= m.activeRange.Max.Y:
		m.setDeltaY(-1)
	case y <= m.activeRange.Min.Y:
		m.setDeltaY(1)
	default:
		m.incrementAndMaybeChangeY()
	}
}

func (m *randomMover) setDeltaY(delta int) {
	m.deltaY = delta
	m.counterY = 0
}

func (m *randomMover) incrementAndMaybeChangeY() {
	m.counterY++
	if m.counterY >= changeInterval {
		m.deltaY = lo.Ternary(rand.Intn(2) == 0, -1, 1)
		m.counterY = 0
	}
}

func (m *randomMover) moveHorizontally() {
	switch m.deltaX {
	case 1:
		m.mr.MoveDirection(input.RightDirection)
	case -1:
		m.mr.MoveDirection(input.LeftDirection)
	}
}

func (m *randomMover) moveVertically() {
	switch m.deltaY {
	case 1:
		m.mr.MoveDirection(input.DownDirection)
	case -1:
		m.mr.MoveDirection(input.UpDirection)
	}
}
