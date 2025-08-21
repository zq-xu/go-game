package actors

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"

	"github.com/zq-xu/go-game/internal/stages/gaming/actors/entity"
	"github.com/zq-xu/go-game/pkg/metric"
)

const ShipName = "Ship"

type Ship struct {
	*entity.Ship
}

func NewShip() *Ship {
	return &Ship{Ship: entity.NewShip()}
}

func (s *Ship) DrawMetrics(screen *ebiten.Image, dc *metric.DrawConfig) {
	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(dc.X), float64(dc.Y))
	op.ColorScale.ScaleWithColor(dc.Color)

	text.Draw(screen, fmt.Sprintf("%s: X: %.0f\tY: %.0f", ShipName, s.X, s.Y), dc.Face, op)
}
