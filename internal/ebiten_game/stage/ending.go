package stage

import (
	"github.com/ebitenui/ebitenui"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/2d-game/internal/ebiten_game/game"
	"github.com/zq-xu/2d-game/internal/ebiten_game/stage/ending"
)

type EndingStage struct {
	ctx *game.Context

	ui *ebitenui.UI

	status game.StageStatus
}

func NewEndingStage(ctx *game.Context, status game.StageStatus) *EndingStage {
	return &EndingStage{
		ctx: ctx,

		ui: ending.NewEndingUI(ctx, status),

		status: status,
	}
}

func (g *EndingStage) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		g.status = game.SuccessStageStatus
	}

	g.ui.Update()

	return nil
}

func (g *EndingStage) Draw(screen *ebiten.Image) {
	g.ui.Draw(screen)
}

func (g *EndingStage) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.ctx.Resource.Cfg.Layout(outsideWidth, outsideHeight)
}

func (g *EndingStage) GoNextStatus() (bool, Interface) {
	switch g.status {
	case game.SuccessStageStatus:
		return true, NewGamingStage(g.ctx)
	default:
		return false, nil
	}
}
