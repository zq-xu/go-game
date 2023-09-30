package img

import (
	"fmt"
	"path"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"

	"github.com/zq-xu/2d-game/internal/ebiten_game/config"
	"github.com/zq-xu/2d-game/internal/ebiten_game/metric"
)

const ShipName = "Ship"

type Ship struct {
	Image

	SpeedFactor float64
}

func NewShipImg(cfg *config.Config) *Ship {
	p := path.Join(config.Cfg.ImageRootPath, "ship.png")

	img := NewImage(p, &cfg.ScreenConfig)

	img.X = (float64(cfg.ScreenWidth - img.Width)) / 2
	img.Y = float64(cfg.ScreenHeight - img.Height)

	s := &Ship{Image: *img, SpeedFactor: 5}
	return s
}

func (s *Ship) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		s.SetX(s.X - s.SpeedFactor)
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		s.SetX(s.X + s.SpeedFactor)
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		s.SetY(s.Y - s.SpeedFactor)
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		s.SetY(s.Y + s.SpeedFactor)
	}
}

func (s *Ship) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(s.X, s.Y)
	screen.DrawImage(s.Image.Image, op)
}

func (s *Ship) DrawMetrics(screen *ebiten.Image, cfg *metric.DrawConfig) {
	text.Draw(screen, fmt.Sprintf("Y: %.0f\tY: %.0f", s.X, s.Y), cfg.Face, cfg.X, cfg.Y, cfg.Color)
}
