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
	SpeedFactor float64
}

func NewShip(ld *loader.Loader) *Ship {
	entity := graphics.NewImageEntityWithImage(ld.ImageLoader.GetShipImage(), ld.Cfg.ScreenWidth, ld.Cfg.ScreenHeight)

	entity.SetX((float64(ld.Cfg.ScreenWidth - entity.Img.Width)) / 2)
	entity.SetY(float64(ld.Cfg.ScreenHeight - entity.Img.Height))

	return &Ship{
		ImageEntity: *entity,
		SpeedFactor: 3,
	}
}

func (s *Ship) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		s.MoveLeft(s.SpeedFactor)
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		s.MoveRight(s.SpeedFactor)
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		s.MoveUp(s.SpeedFactor)
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		s.MoveDown(s.SpeedFactor)
	}
}

func (s *Ship) DrawMetrics(screen *ebiten.Image, cfg *metric.DrawConfig) {
	text.Draw(screen, fmt.Sprintf("Y: %.0f\tY: %.0f", s.X, s.Y), cfg.Face, cfg.X, cfg.Y, cfg.Color)
}
