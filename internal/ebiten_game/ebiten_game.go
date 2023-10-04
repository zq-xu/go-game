package ebitengame

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/2d-game/internal/ebiten_game/config"
	"github.com/zq-xu/2d-game/internal/ebiten_game/game"
	"github.com/zq-xu/2d-game/internal/ebiten_game/stage"
)

type Game struct {
	ctx   *game.Context
	stage *stage.StageController
}

func NewGame() *Game {
	// ebiten.SetFullscreen(config.Cfg.FullScreen)
	ebiten.SetWindowSize(config.Cfg.ScreenWidth, config.Cfg.ScreenHeight)
	ebiten.SetWindowTitle(config.Cfg.Title)

	// ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	// ebiten.SetScreenClearedEveryFrame(false)
	// ebiten.SetVsyncEnabled(false)

	g := &Game{}
	g.ctx = game.NewContext()
	g.stage = stage.NewStageController(g.ctx)

	return g
}

func (g *Game) Update() error {
	return g.stage.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.stage.Draw(screen)

	// ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %f", ebiten.ActualFPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.stage.Layout(outsideWidth, outsideHeight)
}
