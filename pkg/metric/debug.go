package metric

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	DebugMetricsName = "DebugMetrics"
)

var DebugMetricsInstance = &DebugerMetrics{}

type DebugerMetrics struct{}

func (dm *DebugerMetrics) DrawMetrics(screen *ebiten.Image, cfg *DrawConfig) {
	// ebitenutil.DebugPrintAt(screen, fmt.Sprintf("FPS: %f", ebiten.ActualFPS()), cfg.X, cfg.Y)

	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(cfg.X), float64(cfg.Y))
	op.ColorScale.ScaleWithColor(cfg.Color)

	text.Draw(screen, fmt.Sprintf("FPS: %f", ebiten.ActualFPS()), cfg.Face, op)
}
