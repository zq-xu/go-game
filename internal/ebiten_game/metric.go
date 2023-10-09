package ebitengame

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/2d-game/internal/ebiten_game/event"
	"github.com/zq-xu/2d-game/pkg/metric"
)

type MetricListener struct {
	input *event.Input

	MetricPool *metric.Pool
}

func NewMetricListener() *MetricListener {
	l := &MetricListener{
		MetricPool: metric.NewMetricPool(),
	}

	l.input = event.NewInput()
	l.MetricPool.Register(event.InputName, l.input)

	l.MetricPool.Register(metric.DebugMetricsName, metric.DebugMetricsInstance)
	return l
}

func (l *MetricListener) Update() error {
	l.input.Update()
	return nil
}

func (l *MetricListener) Draw(screen *ebiten.Image) {
	l.MetricPool.DrawMetrics(screen)
}
