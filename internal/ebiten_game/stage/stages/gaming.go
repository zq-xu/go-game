package stages

import (
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/2d-game/internal/ebiten_game/game"
	"github.com/zq-xu/2d-game/internal/ebiten_game/stage/gameui"
)

type GamingStage struct {
	ctx *game.Context

	runtime *gameui.Runtime

	navbar *gameui.Navbar

	Status
}

func NewGamingStage(ctx *game.Context) *GamingStage {
	g := &GamingStage{
		ctx: ctx,
	}

	g.runtime = gameui.NewRuntime(ctx,
		gameui.WithRuntimeCollission(func(b bool) {
			if b {
				g.checkoutNextStage(NewEndingStage(g.ctx, game.FailStatus))
			}
		}),
	)

	g.navbar = gameui.NewNavbar(ctx,
		gameui.WithNavbarSettingButonOpts(
			widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
				g.checkoutNextStage(NewSettingStage(g.ctx, g))
			}),
		))

	return g
}

func (g *GamingStage) Update() error {
	err := g.runtime.Update()
	if err != nil {
		return err
	}

	err = g.navbar.Update()
	if err != nil {
		return err
	}

	return nil
}

func (g *GamingStage) Draw(screen *ebiten.Image) {
	screen.Fill(g.ctx.Resource.Cfg.BgColor)

	g.runtime.Draw(screen)
	g.navbar.Draw(screen)
}

func (g *GamingStage) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.ctx.Resource.Layout(outsideWidth, outsideHeight)
}
