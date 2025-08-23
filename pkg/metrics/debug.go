package metrics

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	DebugMetricsName  = "Debug"
	systemMetricsName = "System"
)

var (
	debugMetricsPool MetricsPool
)

type debugerMetrics struct{}

// GetDebugMetricsPool
func GetDebugMetricsPool() MetricsPool {
	if debugMetricsPool == nil {
		debugMetricsPool = NewMetricPool()
		debugMetricsPool.Register(systemMetricsName, &debugerMetrics{})
	}

	return debugMetricsPool
}

func (dm *debugerMetrics) DrawMetrics(screen *ebiten.Image, cfg *DrawConfig) {
	dm.drawFPS(screen, cfg)

	cfg.Y += MetricLineHeight
	dm.drawTPS(screen, cfg)
}

func (dm *debugerMetrics) drawFPS(screen *ebiten.Image, cfg *DrawConfig) {
	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(cfg.X), float64(cfg.Y))
	op.ColorScale.ScaleWithColor(cfg.Color)

	text.Draw(screen, fmt.Sprintf("FPS: %f", ebiten.ActualFPS()), cfg.Face, op)
}

func (dm *debugerMetrics) drawTPS(screen *ebiten.Image, cfg *DrawConfig) {
	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(cfg.X), float64(cfg.Y))
	op.ColorScale.ScaleWithColor(cfg.Color)

	text.Draw(screen, fmt.Sprintf("TPS: %f", ebiten.ActualTPS()), cfg.Face, op)
}
