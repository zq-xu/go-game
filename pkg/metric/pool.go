package metric

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/bitmapfont/v3"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

const metricLineHeight = 16

var colorSet = []color.Color{color.Black, color.RGBA{0xff, 0, 0, 0xff}, color.RGBA{0, 0, 0xff, 0xff}}

type Interface interface {
	DrawMetrics(screen *ebiten.Image, cfg *DrawConfig)
}

type metricObject struct {
	Interface

	name  string
	index int
}

type DrawConfig struct {
	Face  font.Face
	X     int
	Y     int
	Color color.Color
}

type Pool struct {
	metricSet map[string]metricObject
}

func NewMetricPool() *Pool {
	return &Pool{
		metricSet: make(map[string]metricObject),
	}
}

func (p *Pool) Register(name string, i Interface) {
	p.metricSet[name] = metricObject{
		Interface: i,

		name:  name,
		index: len(p.metricSet),
	}
}

func (p *Pool) DrawMetrics(screen *ebiten.Image) {
	p.drawMetrics(screen, metricLineHeight)
}

func (p *Pool) drawMetrics(screen *ebiten.Image, strartY int) {
	for name, obj := range p.metricSet {
		dc := &DrawConfig{
			Face:  bitmapfont.Face,
			X:     4,
			Y:     obj.index*metricLineHeight + strartY,
			Color: colorSet[obj.index%len(colorSet)],
		}
		text.Draw(screen, fmt.Sprintf("%s: ", name), dc.Face, dc.X, dc.Y, dc.Color)

		dc.X = 6 * (len(name) + 2)
		obj.DrawMetrics(screen, dc)

	}
}
