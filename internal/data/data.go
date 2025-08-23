package data

import (
	"time"

	"github.com/zq-xu/go-game/pkg/metrics"
)

type Data interface {
	AddShotUFO(count int)
	GetShotUFO() int64

	Reset()

	Metrics() metrics.PoolList
}

type gameData struct {
	startTime    time.Time
	endTime      time.Time
	timeDuration time.Duration

	shotUFO int64
	score   int64

	metricsPool metrics.PoolList
}

func NewGameData() Data {
	gd := &gameData{
		metricsPool: metrics.NewMetricsPoolList(),
	}
	gd.metricsPool.Add(metrics.DebugMetricsName, metrics.GetDebugMetricsPool())
	return gd
}

func (gd *gameData) AddShotUFO(count int) {
	gd.shotUFO += int64(count)
}

func (gd *gameData) GetShotUFO() int64 {
	return gd.shotUFO
}

func (gd *gameData) Reset() {
	gd.shotUFO = 0
	gd.score = 0
}

func (gd *gameData) Metrics() metrics.PoolList { return gd.metricsPool }
