package entity

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"

	"github.com/zq-xu/2d-game/internal/ebiten_game/game"
	"github.com/zq-xu/2d-game/internal/ebiten_game/resource/loader"
	"github.com/zq-xu/2d-game/pkg/graphics"
	"github.com/zq-xu/2d-game/pkg/metric"
)

const ShipName = "Ship"

// TODO resize for widows size changes
type Ship struct {
	graphics.ImageEntity
	XSpeedFactor float64
	YSpeedFactor float64
}

func NewShip(ctx *game.Context) *Ship {
	img := ctx.Resource.ImageLoader.ImgLoader.MustGetImage(loader.ShipImgPath)
	entity := graphics.NewImageEntityWithImage(img, ctx.Resource.ScreenWidth, ctx.Resource.ScreenHeight)

	entity.SetX((float64(ctx.Resource.ScreenWidth - entity.Img.Width)) / 2)
	entity.SetY(float64(ctx.Resource.ScreenHeight - entity.Img.Height))

	return &Ship{
		ImageEntity:  *entity,
		XSpeedFactor: ctx.Resource.ShipXSpeedFactor,
		YSpeedFactor: ctx.Resource.ShipYSpeedFactor,
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

func (s *Ship) DrawMetrics(screen *ebiten.Image, dc *metric.DrawConfig) {
	text.Draw(screen, fmt.Sprintf("%s: ", ShipName), dc.Face, dc.X, dc.Y, dc.Color)
	text.Draw(screen,
		fmt.Sprintf("X: %.0f\tY: %.0f", s.X, s.Y),
		dc.Face,
		dc.X+metric.MetricCharWidth*(len(ShipName)+5),
		dc.Y,
		dc.Color,
	)

}
