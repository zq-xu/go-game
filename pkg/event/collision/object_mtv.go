package collision

import (
	"math"

	"github.com/solarlune/resolv"
)

type mtv struct {
	X, Y float64
}

func computeMTV(a, b resolv.IShape) mtv {
	dx1 := b.Bounds().Max.X - a.Bounds().Min.X
	dx2 := b.Bounds().Min.X - a.Bounds().Max.X
	dy1 := b.Bounds().Max.Y - a.Bounds().Min.Y
	dy2 := b.Bounds().Min.Y - a.Bounds().Max.Y

	minX := dx1
	if math.Abs(dx2) < math.Abs(minX) {
		minX = dx2
	}
	minY := dy1
	if math.Abs(dy2) < math.Abs(minY) {
		minY = dy2
	}

	if math.Abs(minX) < math.Abs(minY) {
		return mtv{X: minX, Y: 0}
	}
	return mtv{X: 0, Y: minY}
}
