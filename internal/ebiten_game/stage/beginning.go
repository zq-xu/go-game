package stage

import (
	"github.com/ebitenui/ebitenui"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/2d-game/internal/ebiten_game/game"
	"github.com/zq-xu/2d-game/internal/ebiten_game/stage/beginning"
)

type BeginningStage struct {
	ctx *game.Context

	ui     *ebitenui.UI
	status game.StageStatus
}

func NewBeginningStage(ctx *game.Context) *BeginningStage {
	return &BeginningStage{
		ctx: ctx,
		ui:  beginning.NewBeginningUI(ctx),
	}
}

func (g *BeginningStage) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		g.status = game.SuccessStageStatus
	}

	g.ui.Update()

	return nil
}

func (g *BeginningStage) Draw(screen *ebiten.Image) {
	g.ui.Draw(screen)
}

func (g *BeginningStage) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.ctx.Resource.Cfg.Layout(outsideWidth, outsideHeight)
}

func (g *BeginningStage) GoNextStatus() (bool, Interface) {
	switch g.status {
	case game.SuccessStageStatus:
		return true, NewGamingStage(g.ctx)
	default:
		return false, nil
	}
}
