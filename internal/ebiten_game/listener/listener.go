package listener

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/internal/ebiten_game/event"
	"github.com/zq-xu/go-game/pkg/metric"
)

const listenerName = "Listener"

type Listener interface {
	Update() error
	Draw(screen *ebiten.Image)

	GameDataListener() GameDataListener
}

type listener struct {
	MetricPool *metric.Pool

	inputListener *inputListener

	gameDataListener GameDataListener
}

func NewListener() Listener {
	l := &listener{
		MetricPool:       metric.NewMetricPool(),
		inputListener:    NewInputListener(),
		gameDataListener: NewGameDataListener(),
	}

	l.MetricPool.Register(event.InputName, l.inputListener.input)
	l.MetricPool.Register(metric.DebugMetricsName, metric.DebugMetricsInstance)

	metric.MultiPool.Add(listenerName, l.MetricPool)

	return l
}

func (l *listener) Update() error {
	_ = l.inputListener.Update()
	return nil
}

func (l *listener) Draw(screen *ebiten.Image) {
	l.MetricPool.DrawMetrics(screen)
}

func (l *listener) GameDataListener() GameDataListener { return l.gameDataListener }
