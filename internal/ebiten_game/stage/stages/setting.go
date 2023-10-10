package stages

import (
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/2d-game/internal/ebiten_game/game"
	"github.com/zq-xu/2d-game/internal/ebiten_game/stage/gameui"
)

type SettingStage struct {
	ctx *game.Context
	ui  *gameui.Setting

	Status
}

func NewSettingStage(ctx *game.Context, preStage Interface) *SettingStage {
	g := &SettingStage{
		ctx: ctx,
		Status: Status{
			preStage: preStage,
		},
	}

	g.ui = gameui.NewSetting(ctx,
		gameui.WithSettingBackButonOpts(
			widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
				g.checkoutPreStage()
			}),
		))

	return g
}

func (g *SettingStage) Update() error {
	g.ui.Update()

	return nil
}

func (g *SettingStage) Draw(screen *ebiten.Image) {
	g.ui.Draw(screen)
}

func (g *SettingStage) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.ctx.Resource.Layout(outsideWidth, outsideHeight)
}
