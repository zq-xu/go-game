package stages

import (
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/internal/ebiten_game/game"
)

var (
	SuccessText = "SUCCESS"
	FailureText = "FAILED"

	resultTextSet = map[game.Status]string{
		game.SuccessStatus: SuccessText,
		game.FailStatus:    FailureText,
	}
)

type EndingStage struct {
	ctx *game.Context

	ui *ebitenui.UI

	Status
}

func NewEndingStage(ctx *game.Context, status game.Status) *EndingStage {
	return &EndingStage{
		ctx: ctx,
		ui:  NewEndingUI(ctx, status),
	}
}

func (g *EndingStage) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		g.checkoutNextStage(NewGamingStage(g.ctx))
	}

	g.ui.Update()

	return nil
}

func (g *EndingStage) Draw(screen *ebiten.Image) {
	g.ui.Draw(screen)
}

func (g *EndingStage) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.ctx.Resource.Layout(outsideWidth, outsideHeight)
}

func NewEndingUI(ctx *game.Context, status game.Status) *ebitenui.UI {
	titleTexts := []string{"GAME OVER", resultTextSet[status]}

	root := ctx.Resource.LayoutResource.NewCenterRowLayout(400, 10, nil, func(c *widget.Container) {
		for _, v := range titleTexts {
			c.AddChild(ctx.Resource.TextResource.NewCenterText(v, ctx.Resource.FontLoader.TitleFace()))
		}

		for _, v := range ctx.Resource.Cfg.StartHintTexts {
			c.AddChild(ctx.Resource.TextResource.NewCenterText(v, ctx.Resource.FontLoader.Face()))
		}
	})

	return &ebitenui.UI{Container: root}
}
