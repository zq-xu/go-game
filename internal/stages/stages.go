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

	StageName() StageName

	// Reset is used to reset all the data of the stage.
	Reset()

	// Reload is invoked when the game stage is switched as the current running stage
	Reload()
}
