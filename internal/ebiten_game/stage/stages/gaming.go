package stages

import (
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/internal/ebiten_game/game"
	"github.com/zq-xu/go-game/internal/ebiten_game/stage/gameui"
	"github.com/zq-xu/go-game/internal/ebiten_game/stage/listener"
)

type GamingStage struct {
	ctx *game.Context

	runtime *gameui.Runtime

	navbar *gameui.Navbar

	metric *gameui.Metric

	inputLitener *listener.InputListener

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

	g.inputLitener = listener.NewInputListener(ctx, func() bool {
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			g.checkoutNextStage(NewPauseStage(g.ctx, g))
			return true
		}

		return false
	})

	g.metric = gameui.NewMetric(ctx)

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

	g.inputLitener.Update()
	return nil
}

func (g *GamingStage) Draw(screen *ebiten.Image) {
	screen.Fill(g.ctx.Resource.Cfg.BgColor)

	g.runtime.Draw(screen)
	g.navbar.Draw(screen)
	g.metric.Draw(screen)
}

func (g *GamingStage) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.ctx.Resource.Layout(outsideWidth, outsideHeight)
}
