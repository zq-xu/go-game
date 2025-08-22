package entity

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/internal/settings"
	"github.com/zq-xu/go-game/pkg/graphics"
)

const ShipImgPath = "images/ship.png"

// TODO resize for widows size changes
type Ship struct {
	ImageEntity

	XSpeedFactor float64
	YSpeedFactor float64
}

func NewShip() *Ship {
	entity := NewImageEntityWithImage(graphics.GetImage(ShipImgPath),
		settings.GetSettings().ScreenWidth(),
		settings.GetSettings().ScreenHeight())

	entity.SetX((float64(settings.GetSettings().ScreenWidth() - entity.Img.Width())) / 2)
	entity.SetY(float64(settings.GetSettings().ScreenHeight() - entity.Img.Height()))

	return &Ship{
		ImageEntity:  *entity,
		XSpeedFactor: settings.GetSettings().ShipXSpeedFactor(),
		YSpeedFactor: settings.GetSettings().ShipYSpeedFactor(),
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
