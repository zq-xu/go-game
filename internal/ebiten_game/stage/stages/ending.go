package stages

import (
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/2d-game/internal/ebiten_game/game"
)

var (
	SuccessText = "SUCCESS"
	FailureText = "FAILED"

	resultTextSet = map[game.StageStatus]string{
		game.SuccessStageStatus: SuccessText,
		game.FailStageStatus:    FailureText,
	}
)

type EndingStage struct {
	ctx *game.Context

	ui *ebitenui.UI

	status game.StageStatus
}

func NewEndingStage(ctx *game.Context, status game.StageStatus) *EndingStage {
	return &EndingStage{
		ctx: ctx,

		ui: NewEndingUI(ctx, status),

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
	return g.ctx.Resource.Layout(outsideWidth, outsideHeight)
}

func (g *EndingStage) GoNextStatus() (bool, Interface) {
	switch g.status {
	case game.SuccessStageStatus:
		return true, NewGamingStage(g.ctx)
	default:
		return false, nil
	}
}

func NewEndingUI(ctx *game.Context, status game.StageStatus) *ebitenui.UI {
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
