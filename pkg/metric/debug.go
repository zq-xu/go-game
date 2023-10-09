package metric

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	DebugMetricsName = "DebugMetrics"
)

var DebugMetricsInstance = &DebugerMetrics{}

type DebugerMetrics struct{}

func (dm *DebugerMetrics) DrawMetrics(screen *ebiten.Image, cfg *DrawConfig) {
	// ebitenutil.DebugPrintAt(screen, fmt.Sprintf("FPS: %f", ebiten.ActualFPS()), cfg.X, cfg.Y)
	text.Draw(screen, fmt.Sprintf("FPS: %f", ebiten.ActualFPS()), cfg.Face, cfg.X, cfg.Y, cfg.Color)
}
