package stages

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	GamingStage    StageName = "gaming"
	BeginningStage StageName = "beginning"
	EndingStage    StageName = "ending"
	MenuStage      StageName = "menu"
	PauseStage     StageName = "pause"
	SettingStage   StageName = "setting"
)

type StageName string

type GameStage interface {
	ebiten.Game

	NextGameStage() GameStage
	Reload()
}
