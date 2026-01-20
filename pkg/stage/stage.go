package stage

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type StageName string

type Stage interface {
	ebiten.Game

	StageName() StageName
}
