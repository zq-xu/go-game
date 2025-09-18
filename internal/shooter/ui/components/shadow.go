package components

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func GenerateShadowDrawerFn(scale float32) func(screen *ebiten.Image) {
	var shadowImage *ebiten.Image
	var op *ebiten.DrawImageOptions

	if scale == 0 {
		scale = 0.2
	}

	return func(screen *ebiten.Image) {
		if shadowImage == nil {
			shadowImage = ebiten.NewImage(screen.Bounds().Dx(), screen.Bounds().Dy())
			shadowImage.Fill(color.White)
		}

		if op == nil {
			op = &ebiten.DrawImageOptions{}
			op.ColorScale.ScaleAlpha(scale)
		}

		screen.DrawImage(shadowImage, op)
	}
}
