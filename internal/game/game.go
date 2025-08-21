package game

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/internal/config"
	"github.com/zq-xu/go-game/internal/listener"
)

type Game struct {
	ebiten.Game

	listener listener.Listener
}

func NewGame() *Game {
	// ebiten.SetFullscreen(config.Cfg.FullScreen)
	// ebiten.SetScreenClearedEveryFrame(false)
	// ebiten.SetVsyncEnabled(false)

	ebiten.SetWindowSize(config.Cfg.ScreenConfig.ScreenWidth, config.Cfg.ScreenConfig.ScreenHeight)
	ebiten.SetWindowTitle(config.Cfg.Title)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	return &Game{
		Game:     NewStageController(),
		listener: listener.GetListener(),
	}
}

func (g *Game) Update() error {
	g.listener.Update()
	return g.Game.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.listener.Draw(screen)
	g.Game.Draw(screen)
}
