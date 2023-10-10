package ebitengame

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/2d-game/internal/ebiten_game/config"
	"github.com/zq-xu/2d-game/internal/ebiten_game/game"
	"github.com/zq-xu/2d-game/internal/ebiten_game/stage"
	"github.com/zq-xu/2d-game/pkg/metric"
)

type Game struct {
	ctx *game.Context

	stage *stage.StageController

	gmListener *GlobalMetricListener
}

func NewGame() *Game {
	// ebiten.SetFullscreen(config.Cfg.FullScreen)
	ebiten.SetWindowSize(config.Cfg.ScreenConfig.ScreenWidth, config.Cfg.ScreenConfig.ScreenHeight)
	ebiten.SetWindowTitle(config.Cfg.Title)

	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	// ebiten.SetScreenClearedEveryFrame(false)
	// ebiten.SetVsyncEnabled(false)

	g := &Game{}
	g.ctx = game.NewContext()

	g.stage = stage.NewStageController(g.ctx)
	g.gmListener = NewGlobalMetricListener()

	return g
}

func (g *Game) Update() error {
	// g.ctx.Resource.Listen()

	g.gmListener.Update()
	return g.stage.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.stage.Draw(screen)

	metric.MultiPool.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.stage.Layout(outsideWidth, outsideHeight)
}
