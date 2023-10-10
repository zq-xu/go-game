package stage

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/2d-game/internal/ebiten_game/event"
	"github.com/zq-xu/2d-game/internal/ebiten_game/game"
	"github.com/zq-xu/2d-game/internal/ebiten_game/stage/stages"
)

type StageController struct {
	ctx *game.Context

	Input *event.Input

	preStage     stages.Interface
	currentStage stages.Interface
}

func NewStageController(ctx *game.Context) *StageController {
	g := &StageController{}

	g.ctx = ctx
	g.currentStage = stages.NewMenuStage(ctx)

	return g
}

func (g *StageController) Update() error {
	b, i := g.currentStage.GoNextStatus()
	if b {
		g.preStage = g.currentStage
		g.currentStage = i
	}

	err := g.currentStage.Update()
	if err != nil {
		return err
	}

	return nil
}

func (g *StageController) Draw(screen *ebiten.Image) {
	g.currentStage.Draw(screen)
}

func (g *StageController) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.currentStage.Layout(outsideWidth, outsideHeight)
}
