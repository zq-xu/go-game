package stages

import "github.com/hajimehoshi/ebiten/v2"

type TempDrawer interface {
	Draw(screen *ebiten.Image)
}

type defaultTempDrawer struct{}

func newDefaultTempDrawer() TempDrawer { return &defaultTempDrawer{} }

func (d *defaultTempDrawer) Draw(screen *ebiten.Image) {}
