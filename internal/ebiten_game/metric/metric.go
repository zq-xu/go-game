package metric

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/bitmapfont/v3"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

var metricSet = make(map[string]metricObject)

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

func Register(name string, i Interface) {
	_, ok := metricSet[name]
	if ok {
		return
	}

	metricSet[name] = metricObject{
		Interface: i,

		name:  name,
		index: len(metricSet),
	}
}

func DrawMetrics(screen *ebiten.Image) {
	for name, obj := range metricSet {
		dc := &DrawConfig{
			Face:  bitmapfont.Face,
			X:     4,
			Y:     (obj.index + 1) * 16,
			Color: colorSet[obj.index%len(colorSet)],
		}
		text.Draw(screen, fmt.Sprintf("%s: ", name), dc.Face, dc.X, dc.Y, dc.Color)

		dc.X = 6 * (len(name) + 2)
		obj.DrawMetrics(screen, dc)

	}
}
