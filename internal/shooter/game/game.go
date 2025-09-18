package game

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/internal/shooter/data"
	"github.com/zq-xu/go-game/pkg/config"
)

type Game struct {
	ebiten.Game

	data data.Data
}

func NewGame() *Game {
	// ebiten.SetFullscreen(config.Cfg.FullScreen)
	// ebiten.SetScreenClearedEveryFrame(false)
	// ebiten.SetVsyncEnabled(false)

	ebiten.SetWindowSize(config.Cfg.ScreenConfig.ScreenWidth, config.Cfg.ScreenConfig.ScreenHeight)
	ebiten.SetWindowTitle(config.Cfg.Title)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	gameData := data.NewGameData()
	return &Game{
		data: gameData,
		Game: NewStageController(gameData),
	}
}
