package entity

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"

	"github.com/zq-xu/2d-game/internal/ebiten_game/loader"
	"github.com/zq-xu/2d-game/pkg/graphics"
	"github.com/zq-xu/2d-game/pkg/metric"
)

const ShipName = "Ship"

type Ship struct {
	graphics.ImageEntity
	XSpeedFactor float64
	YSpeedFactor float64
}

func NewShip(ld *loader.Loader) *Ship {
	entity := graphics.NewImageEntityWithImage(ld.ImageLoader.GetShipImage(), ld.Cfg.ScreenWidth, ld.Cfg.ScreenHeight)

	entity.SetX((float64(ld.Cfg.ScreenWidth - entity.Img.Width)) / 2)
	entity.SetY(float64(ld.Cfg.ScreenHeight - entity.Img.Height))

	return &Ship{
		ImageEntity:  *entity,
		XSpeedFactor: ld.Cfg.ShipXSpeedFactor,
		YSpeedFactor: ld.Cfg.ShipYSpeedFactor,
	}
}

func (s *Ship) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		s.MoveLeft(s.XSpeedFactor)
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		s.MoveRight(s.XSpeedFactor)
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		s.MoveUp(s.YSpeedFactor)
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		s.MoveDown(s.YSpeedFactor)
	}
}

func (s *Ship) DrawMetrics(screen *ebiten.Image, cfg *metric.DrawConfig) {
	text.Draw(screen, fmt.Sprintf("X: %.0f\tY: %.0f", s.X, s.Y), cfg.Face, cfg.X, cfg.Y, cfg.Color)
}
