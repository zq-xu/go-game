package stage

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"

	"github.com/zq-xu/2d-game/internal/ebiten_game/config"
	"github.com/zq-xu/2d-game/internal/ebiten_game/loader"
)

var (
	TitleTexts = []string{config.Cfg.Title, "AUTHOR: ZHIQIANG XU"}
	StartTexts = []string{"", "", "", "", "PRESS ENTER KEY TO START"}
)

type BeginningStage struct {
	loader *loader.Loader
	status StageStatus
}

func NewBeginningStage(loader *loader.Loader) *BeginningStage {
	return &BeginningStage{
		loader: loader,
	}
}

func (g *BeginningStage) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		g.status = SuccessStageStatus
	}

	return nil
}

func (g *BeginningStage) Draw(screen *ebiten.Image) {
	tittleHight := len(TitleTexts) * 2 * g.loader.Cfg.TitleFontSize
	textHeight := len(StartTexts) * 2 * g.loader.Cfg.FontSize

	startY := (g.loader.Cfg.ScreenHeight - tittleHight - textHeight) / 2

	for _, l := range TitleTexts {
		x := (g.loader.Cfg.ScreenWidth - len(l)*g.loader.Cfg.TitleFontSize) / 2
		startY += 2 * g.loader.Cfg.TitleFontSize
		text.Draw(screen, l, g.loader.FontLoader.TitleArcadeFont, x, startY, color.White)
	}

	for _, l := range StartTexts {
		x := (g.loader.Cfg.ScreenWidth - len(l)*g.loader.Cfg.FontSize) / 2
		startY += 2 * g.loader.Cfg.FontSize
		text.Draw(screen, l, g.loader.FontLoader.ArcadeFont, x, startY, color.White)
	}
}

func (g *BeginningStage) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.loader.Cfg.Layout(outsideWidth, outsideHeight)
}

func (g *BeginningStage) GoNextStatus() (bool, Interface) {
	switch g.status {
	case SuccessStageStatus:
		return true, NewGamingStage(g.loader)
	default:
		return false, nil
	}
}
