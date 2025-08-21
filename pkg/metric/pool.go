package metric

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var colorSet = []color.Color{color.RGBA{0xff, 0, 0, 0xff}, color.RGBA{0, 0xff, 0, 0xff}, color.RGBA{0, 0, 0xff, 0xff}}

type Interface interface {
	DrawMetrics(screen *ebiten.Image, cfg *DrawConfig)
}

type metricObject struct {
	Interface

	name  string
	index int
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
		name:      name,
		index:     len(p.metricSet),
	}
}

func (p *Pool) DrawMetrics(screen *ebiten.Image) {
	p.drawMetrics(screen, defaultStartX, MetricLineHeight)
}

func (p *Pool) drawMetrics(screen *ebiten.Image, strartX, strartY int) {
	for _, obj := range p.metricSet {
		dc := DefaultDrawConfig.Copy()

		dc.X = strartX
		dc.Y = obj.index*MetricLineHeight + strartY
		dc.Color = colorSet[obj.index%len(colorSet)]

		obj.DrawMetrics(screen, dc)
	}
}
