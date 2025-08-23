package metrics

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var colorSet = []color.Color{color.RGBA{0xff, 0, 0, 0xff}, color.RGBA{0, 0xff, 0, 0xff}, color.RGBA{0, 0, 0xff, 0xff}}

type Metrics interface {
	DrawMetrics(screen *ebiten.Image, cfg *DrawConfig)
}

type MetricsPool interface {
	Register(name string, i Metrics)
	DrawMetrics(screen *ebiten.Image, startX, startY int)

	Length() int
}

type metricObject struct {
	Metrics

	name string
}

type metricPool struct {
	metricSet []metricObject
}

func NewMetricPool() MetricsPool {
	return &metricPool{
		metricSet: make([]metricObject, 0),
	}
}

func (p *metricPool) Register(name string, i Metrics) {
	p.metricSet = append(p.metricSet, metricObject{
		Metrics: i,
		name:    name,
	})
}

func (p *metricPool) Length() int {
	return len(p.metricSet)
}

func (p *metricPool) DrawMetrics(screen *ebiten.Image, startX, startY int) {
	for index, obj := range p.metricSet {
		p.drawItemTitle(screen, obj.name, startX, startY)
		startY += MetricLineHeight
		p.drawItemContent(screen, obj, startX, startY, colorSet[index%len(colorSet)])
		startY += MetricLineHeight
	}
}

func (pi *metricPool) drawItemTitle(screen *ebiten.Image, txt string, startX, startY int) {
	dc := DefaultDrawConfig.Copy()

	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(startX), float64(startY))
	op.ColorScale.ScaleWithColor(dc.Color)

	text.Draw(screen, fmt.Sprintf("%s: ", txt), dc.Face, op)
}

func (pi *metricPool) drawItemContent(screen *ebiten.Image, obj Metrics, startX, startY int, clr color.Color) {
	dc := DefaultDrawConfig.Copy()
	dc.X = startX + MetricCharWidth
	dc.Y = startY
	dc.Color = clr

	obj.DrawMetrics(screen, dc)
}
