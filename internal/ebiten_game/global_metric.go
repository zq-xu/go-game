package ebitengame

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/internal/ebiten_game/event"
	"github.com/zq-xu/go-game/pkg/metric"
)

const globalMetricName = "Global"

type GlobalMetricListener struct {
	input *event.Input

	MetricPool *metric.Pool
}

func NewGlobalMetricListener() *GlobalMetricListener {
	l := &GlobalMetricListener{}

	l.MetricPool = metric.NewMetricPool()
	metric.MultiPool.Add(globalMetricName, l.MetricPool)

	l.input = event.NewInput()
	l.MetricPool.Register(event.InputName, l.input)
	l.MetricPool.Register(metric.DebugMetricsName, metric.DebugMetricsInstance)

	return l
}

func (l *GlobalMetricListener) Update() error {
	l.input.Update()
	return nil
}

func (l *GlobalMetricListener) Draw(screen *ebiten.Image) {
	l.MetricPool.DrawMetrics(screen)
}
