package listener

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type metricListener struct{}

func NewMetricListener() *metricListener {
	return &metricListener{}
}

func (l *metricListener) Update() error { return nil }

func (l *metricListener) Draw(screen *ebiten.Image) {
}
