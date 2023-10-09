package stages

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/2d-game/internal/ebiten_game/game"
	"github.com/zq-xu/2d-game/internal/ebiten_game/stage/gameui"
)

type MenuStage struct {
	ctx *game.Context
	ui  *gameui.MainMenu
}

func NewMenuStage(ctx *game.Context) *MenuStage {
	return &MenuStage{
		ctx: ctx,
		ui:  gameui.NewMainMenu(ctx),
	}
}

func (g *MenuStage) Update() error {
	g.ui.Update()
	return nil
}

func (g *MenuStage) Draw(screen *ebiten.Image) {
	g.ui.Draw(screen)
}

func (g *MenuStage) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.ctx.Resource.Layout(outsideWidth, outsideHeight)

}

func (g *MenuStage) GoNextStatus() (bool, Interface) {
	switch g.ui.Status {
	case gameui.StartMenuStatus:
		return true, NewGamingStage(g.ctx)
	case gameui.BackMenuStatus:
		return true, NewBeginningStage(g.ctx)
	}

	return false, nil
}
