package stages

import (
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/2d-game/internal/ebiten_game/game"
)

type BeginningStage struct {
	ctx *game.Context

	ui     *ebitenui.UI
	status game.StageStatus
}

func NewBeginningStage(ctx *game.Context) *BeginningStage {
	return &BeginningStage{
		ctx: ctx,
		ui:  NewBeginningUI(ctx),
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
	return g.ctx.Resource.Layout(outsideWidth, outsideHeight)
}

func (g *BeginningStage) GoNextStatus() (bool, Interface) {
	switch g.status {
	case game.SuccessStageStatus:
		return true, NewGamingStage(g.ctx)
	default:
		return false, nil
	}
}

func NewBeginningUI(ctx *game.Context) *ebitenui.UI {
	root := ctx.Resource.LayoutResource.NewCenterRowLayout(400, 10, nil, func(c *widget.Container) {
		c.AddChild(ctx.Resource.TextResource.NewCenterText(ctx.Resource.Cfg.Title, ctx.Resource.FontLoader.TitleFace()))
		c.AddChild(ctx.Resource.TextResource.NewCenterText(ctx.Resource.Cfg.AuthorText, ctx.Resource.FontLoader.TitleFace()))

		for _, v := range ctx.Resource.Cfg.StartHintTexts {
			c.AddChild(ctx.Resource.TextResource.NewCenterText(v, ctx.Resource.FontLoader.Face()))
		}
	})

	return &ebitenui.UI{Container: root}
}
