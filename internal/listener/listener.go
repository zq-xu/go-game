package listener

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/pkg/metric"
)

const listenerName = "Listener"

var gloableListener Listener

type Listener interface {
	Update() error
	Draw(screen *ebiten.Image)

	GameDataListener() GameDataListener
}

type listener struct {
	MetricPool       *metric.Pool
	gameDataListener GameDataListener
}

func GetListener() Listener {
	if gloableListener == nil {
		gloableListener = newListener()
	}
	return gloableListener
}

func newListener() Listener {
	l := &listener{
		MetricPool:       metric.NewMetricPool(),
		gameDataListener: NewGameDataListener(),
	}

	l.MetricPool.Register(metric.DebugMetricsName, metric.DebugMetricsInstance)
	metric.MultiPool.Add(listenerName, l.MetricPool)

	return l
}

func (l *listener) Update() error {
	return nil
}

func (l *listener) Draw(screen *ebiten.Image) {
	l.MetricPool.DrawMetrics(screen)
}

func (l *listener) GameDataListener() GameDataListener { return l.gameDataListener }
