package background

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/internal/settings"
)

type yGradientUpwards struct {
	yGradientBackground
}

func NewYGradientUpwards(bg *ebiten.Image, bgW, bgH int) ebiten.Game {
	return &yGradientUpwards{yGradientBackground: *newYGradientBackground(bg, bgW, bgH)}
}

func (g *yGradientUpwards) Update() error {
	// The background scrolls down (the character appears to fly/run upwards)
	g.offsetY += speed

	// Keep offsetY within the height range of the image.
	if g.offsetY > float64(g.bgH) {
		g.offsetY -= float64(g.bgH)
	}
	return nil
}

func (g *yGradientUpwards) Draw(screen *ebiten.Image) {
	// Starting from -offset, draw horizontally tiled downwards.
	startY := -math.Mod(g.offsetY, float64(g.bgH))

	for y := startY; y < float64(settings.GetSettings().ScreenHeight()); y += float64(g.bgH) {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(0, y)
		screen.DrawImage(g.bg, op)
	}
}

func (g *yGradientUpwards) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func NewDefaultYGradientUpwards(w, h int) ebiten.Game {
	bg := newGradientBackgroundImage(w, h)
	return NewYGradientUpwards(bg, w, h)
}
