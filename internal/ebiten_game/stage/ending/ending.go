package ending

import (
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"

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

func NewEndingUI(ctx *game.Context, status game.StageStatus) *ebitenui.UI {
	titleTexts := []string{"GAME OVER", resultTextSet[status]}

	root := ctx.Resource.LayoutResource.NewCenterLayout(400, 10, nil, func(c *widget.Container) {
		for _, v := range titleTexts {
			c.AddChild(ctx.Resource.TextResource.NewCenterText(v, ctx.Resource.FontLoader.TitleFace()))
		}

		for _, v := range ctx.Resource.Cfg.StartHintTexts {
			c.AddChild(ctx.Resource.TextResource.NewCenterText(v, ctx.Resource.FontLoader.Face()))
		}
	})

	return &ebitenui.UI{Container: root}
}
