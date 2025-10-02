package entity

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rotisserie/eris"

	"github.com/zq-xu/go-game/assets"
	"github.com/zq-xu/go-game/internal/shooter/settings"
)

var ShipImgPath = assets.GetShooterImagePath("ship.png")

// TODO resize for widows size changes
type Ship struct {
	ImageEntity

	XSpeedFactor float64
	YSpeedFactor float64
}

func NewShip() (*Ship, error) {
	entity, err := NewImageEntity(
		ShipImgPath,
		settings.GetSettings().ScreenWidth(),
		settings.GetSettings().ScreenHeight())
	if err != nil {
		return nil, eris.Wrap(err, "new ship image entity failed")
	}

	fmt.Println("ship width", entity.Img.Width())
	fmt.Println("ship height", entity.Img.Height())

	entity.SetX((float64(settings.GetSettings().ScreenWidth() - entity.Img.Width())) / 2)
	entity.SetY(float64(settings.GetSettings().ScreenHeight() - entity.Img.Height()))

	return &Ship{
		ImageEntity:  *entity,
		XSpeedFactor: settings.GetSettings().ShipXSpeedFactor(),
		YSpeedFactor: settings.GetSettings().ShipYSpeedFactor(),
	}, nil
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
