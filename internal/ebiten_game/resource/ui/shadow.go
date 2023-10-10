package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/2d-game/internal/ebiten_game/config"
)

type ShadowResource struct {
	cfg *config.Config

	shadowImage *ebiten.Image

	op *ebiten.DrawImageOptions
}

func NewShadowResource(cfg *config.Config) *ShadowResource {
	return &ShadowResource{cfg: cfg}
}

func (s *ShadowResource) DrawShadow(screen *ebiten.Image) {
	if s.shadowImage == nil {
		s.shadowImage = ebiten.NewImage(s.cfg.ScreenConfig.ScreenWidth, s.cfg.ScreenConfig.ScreenHeight)
		s.shadowImage.Fill(color.Black)

	}

	if s.op == nil {
		s.op = &ebiten.DrawImageOptions{}
		s.op.ColorScale.ScaleAlpha(0.7)
	}

	screen.DrawImage(s.shadowImage, s.op)
}
