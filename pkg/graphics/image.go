package graphics

import "github.com/hajimehoshi/ebiten/v2"

// DrawImage
func DrawImage(screen, img *ebiten.Image, x, y float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	screen.DrawImage(img, op)
}
