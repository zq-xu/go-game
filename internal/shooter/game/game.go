package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rotisserie/eris"

	"github.com/zq-xu/go-game/internal/shooter/data"
	"github.com/zq-xu/go-game/pkg/config"
)

type Game struct {
	ebiten.Game

	data data.Data
}

// NewGame
func NewGame() (*Game, error) {
	// ebiten.SetFullscreen(config.Cfg.FullScreen)
	// ebiten.SetScreenClearedEveryFrame(false)
	// ebiten.SetVsyncEnabled(false)

	ebiten.SetWindowSize(config.Cfg.ScreenConfig.ScreenWidth, config.Cfg.ScreenConfig.ScreenHeight)
	ebiten.SetWindowTitle(config.Cfg.Title)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	gameData := data.NewGameData()
	g, err := NewStageController(gameData)
	if err != nil {
		return nil, eris.Wrap(err, "new stage controller failed.")
	}

	return &Game{
		data: gameData,
		Game: g,
	}, nil
}
