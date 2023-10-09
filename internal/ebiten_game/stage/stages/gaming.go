package stages

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/2d-game/internal/ebiten_game/game"
	"github.com/zq-xu/2d-game/internal/ebiten_game/stage/gameui"
)

type GamingStage struct {
	ctx *game.Context

	runtime *gameui.Runtime
	status  game.StageStatus

	navbar *gameui.Navbar
}

func NewGamingStage(ctx *game.Context) *GamingStage {
	return &GamingStage{
		ctx:     ctx,
		runtime: gameui.NewRuntime(ctx),
		navbar:  gameui.NewNavbar(ctx),
	}
}

func (g *GamingStage) Update() error {

	err := g.runtime.Update()
	if err != nil {
		return err
	}

	g.navbar.Update()
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

func (g *GamingStage) GoNextStatus() (bool, Interface) {
	switch g.runtime.GetStatus() {
	case game.SuccessStageStatus, game.FailStageStatus:
		return true, NewEndingStage(g.ctx, g.status)
	default:
		return false, nil
	}
}
