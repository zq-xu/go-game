package ebitengame

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/2d-game/internal/ebiten_game/config"
	"github.com/zq-xu/2d-game/internal/ebiten_game/data"
	"github.com/zq-xu/2d-game/internal/ebiten_game/mode"
)

type Game struct {
	data *data.GameData

	modeHandler mode.GameModeHandler
}

func NewGame() *Game {
	ebiten.SetWindowSize(config.Cfg.ScreenWidth, config.Cfg.ScreenHeight)
	ebiten.SetWindowTitle(config.Cfg.Title)

	g := &Game{
		data: data.NewGameData(),
	}

	g.SetMode(mode.WaitingStartMode)
	return g
}

// SetMode: implement ModeListener
func (g *Game) SetMode(m mode.Mode) {
	g.modeHandler, _ = mode.NewModeHandler(m, g.data, g)
}

func (g *Game) Update() error {
	return g.modeHandler.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.modeHandler.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.modeHandler.Layout(outsideWidth, outsideHeight)
}
