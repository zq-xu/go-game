package beginning

import (
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"

	"github.com/zq-xu/2d-game/internal/ebiten_game/game"
)

func NewBeginningUI(ctx *game.Context) *ebitenui.UI {
	root := ctx.Resource.LayoutResource.NewCenterLayout(400, 10, nil, func(c *widget.Container) {
		c.AddChild(ctx.Resource.TextResource.NewCenterText(ctx.Resource.Cfg.Title, ctx.Resource.FontLoader.TitleFace()))
		c.AddChild(ctx.Resource.TextResource.NewCenterText(ctx.Resource.Cfg.AuthorText, ctx.Resource.FontLoader.TitleFace()))

		for _, v := range ctx.Resource.Cfg.StartHintTexts {
			c.AddChild(ctx.Resource.TextResource.NewCenterText(v, ctx.Resource.FontLoader.Face()))
		}
	})

	return &ebitenui.UI{Container: root}
}
