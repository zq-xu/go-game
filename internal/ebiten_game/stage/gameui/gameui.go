package gameui

import "github.com/hajimehoshi/ebiten/v2"

type Interface interface {
	ebiten.Game

	GoNextStatus() (bool, Interface)
}
