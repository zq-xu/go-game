package ebitengame

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/2d-game/internal/ebiten_game/config"
	"github.com/zq-xu/2d-game/internal/ebiten_game/loader"
	"github.com/zq-xu/2d-game/internal/ebiten_game/stage"
)

type Game struct {
	loader *loader.Loader
	stage  *stage.StageController
}

func NewGame() *Game {
	ebiten.SetFullscreen(config.Cfg.FullScreen)
	ebiten.SetWindowSize(config.Cfg.ScreenWidth, config.Cfg.ScreenHeight)
	ebiten.SetWindowTitle(config.Cfg.Title)

	g := &Game{}
	g.loader = loader.NewLoader()
	g.stage = stage.NewStageController(g.loader)

	return g
}

func (g *Game) Update() error {
	return g.stage.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.stage.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.stage.Layout(outsideWidth, outsideHeight)
}
