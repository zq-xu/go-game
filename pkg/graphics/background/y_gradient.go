package background

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	speed = 2.0 // Number of scrolling pixels per frame (default 60 FPS)
)

type yGradientBackground struct {
	bg       *ebiten.Image
	bgW, bgH int
	offsetY  float64
}

func newYGradientBackground(bg *ebiten.Image, bgW, bgH int) *yGradientBackground {
	return &yGradientBackground{bg: bg, bgW: bgW, bgH: bgH}
}

func newGradientBackgroundImage(w, h int) *ebiten.Image {
	img := ebiten.NewImage(w, h)

	lightBlue := color.RGBA{135, 206, 250, 255}   // SkyBlue
	lightYellow := color.RGBA{255, 255, 153, 255} // LightYellow

	half := h / 2

	for y := 0; y < h; y++ {
		var c color.RGBA
		if y < half {
			// first half: SkyBlue to LightYellow
			t := float64(y) / float64(half-1)
			r := uint8(float64(lightBlue.R)*(1-t) + float64(lightYellow.R)*t)
			g := uint8(float64(lightBlue.G)*(1-t) + float64(lightYellow.G)*t)
			b := uint8(float64(lightBlue.B)*(1-t) + float64(lightYellow.B)*t)
			c = color.RGBA{r, g, b, 255}
		} else {
			// second half: LightYellow to SkyBlue
			t := float64(y-half) / float64(half-1)
			r := uint8(float64(lightYellow.R)*(1-t) + float64(lightBlue.R)*t)
			g := uint8(float64(lightYellow.G)*(1-t) + float64(lightBlue.G)*t)
			b := uint8(float64(lightYellow.B)*(1-t) + float64(lightBlue.B)*t)
			c = color.RGBA{r, g, b, 255}
		}

		for x := 0; x < w; x++ {
			img.Set(x, y, c)
		}
	}
	return img
}
