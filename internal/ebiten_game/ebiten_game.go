package ebitengame

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/internal/ebiten_game/config"
	"github.com/zq-xu/go-game/internal/ebiten_game/game"
	"github.com/zq-xu/go-game/internal/ebiten_game/stage"
)

type Game struct {
	ctx *game.Context

	stage *stage.StageController
}

func NewGame() *Game {
	// ebiten.SetFullscreen(config.Cfg.FullScreen)
	// ebiten.SetScreenClearedEveryFrame(false)
	// ebiten.SetVsyncEnabled(false)

	ebiten.SetWindowSize(config.Cfg.ScreenConfig.ScreenWidth, config.Cfg.ScreenConfig.ScreenHeight)
	ebiten.SetWindowTitle(config.Cfg.Title)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	g := &Game{ctx: game.NewContext()}

	g.stage = stage.NewStageController(g.ctx)

	return g
}

func (g *Game) Update() error {
	// g.ctx.Resource.Listen()

	g.ctx.Listener.Update()
	return g.stage.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.stage.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.stage.Layout(outsideWidth, outsideHeight)
}
