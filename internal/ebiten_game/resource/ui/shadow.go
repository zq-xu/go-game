package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/2d-game/internal/ebiten_game/config"
)

type ShadowResource struct {
	cfg *config.Config
}

func NewShadowResource(cfg *config.Config) *ShadowResource {
	return &ShadowResource{cfg: cfg}
}

// func (s *ShadowResource) DrawShadow(screen *ebiten.Image) {
// 	if s.shadowImage == nil {
// 		s.shadowImage = ebiten.NewImage(screen.Bounds().Dx(), screen.Bounds().Dy())
// 		s.shadowImage.Fill(color.Black)

// 	}

// 	if s.op == nil {
// 		s.op = &ebiten.DrawImageOptions{}
// 		s.op.ColorScale.ScaleAlpha(0.7)
// 	}

// 	screen.DrawImage(s.shadowImage, s.op)
// }

func (s *ShadowResource) GenerateShadowDrawerFn(scale float32) func(screen *ebiten.Image) {
	var shadowImage *ebiten.Image
	var op *ebiten.DrawImageOptions

	if scale == 0 {
		scale = 0.3
	}

	return func(screen *ebiten.Image) {
		if shadowImage == nil {
			shadowImage = ebiten.NewImage(screen.Bounds().Dx(), screen.Bounds().Dy())
			shadowImage.Fill(color.Black)
		}

		if op == nil {
			op = &ebiten.DrawImageOptions{}
			op.ColorScale.ScaleAlpha(scale)
		}

		screen.DrawImage(shadowImage, op)
	}

}
