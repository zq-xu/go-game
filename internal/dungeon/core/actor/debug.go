package actor

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/pkg/graphics"
	"github.com/zq-xu/go-game/pkg/graphics/images/imagekit"
)

// DebugCollisionBorder
func (a *actor) debugCollisionBorder(screen *ebiten.Image) {
	rect := ebiten.NewImage(int(a.op.Width), int(a.op.Height))
	rect.Fill(color.RGBA{255, 0, 0, 128})

	leftTopX, leftTopY := a.LeftTop()
	graphics.DrawImage(screen, rect, leftTopX, leftTopY)
}

// DebugImageBorder 	 debug draw border
func (a *actor) debugImageBorder(screen *ebiten.Image, img imagekit.Image) {
	rect := ebiten.NewImage(img.Width(), img.Height())
	rect.Fill(color.RGBA{0, 0, 255, 128})

	x, y := a.Center()
	leftTopX := x - float64(img.Width())/2
	leftTopY := y - float64(img.Height())/2
	graphics.DrawImage(screen, rect, leftTopX, leftTopY)
}
