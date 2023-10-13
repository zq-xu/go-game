package stages

import (
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/internal/ebiten_game/game"
)

type MenuStage struct {
	ctx *game.Context

	ui *ebitenui.UI
	Status
}

func NewMenuStage(ctx *game.Context) *MenuStage {
	g := &MenuStage{ctx: ctx}

	g.generateMenuUI()

	return g
}

func (g *MenuStage) generateMenuUI() {
	menuWidth := g.ctx.Resource.ScreenWidth / 3
	root := g.ctx.Resource.LayoutResource.NewCenterRowLayout(menuWidth, 10, nil, func(c *widget.Container) {
		c.AddChild(g.ctx.Resource.ButtonResource.NewMenuButton("Start",
			widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
				g.checkoutNextStage(NewBeginningStage(g.ctx))
			})))

		c.AddChild(g.ctx.Resource.SeparatorResource.NewSeparator(widget.RowLayoutData{Stretch: true}))

		c.AddChild(g.ctx.Resource.ButtonResource.NewMenuButton("Next",
			widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
				g.checkoutNextStage(NewGamingStage(g.ctx))
			})))

		c.AddChild(g.ctx.Resource.SeparatorResource.NewSeparator(widget.RowLayoutData{Stretch: true}))

		c.AddChild(g.ctx.Resource.ButtonResource.NewMenuButton("Back"))
	})

	g.ui = &ebitenui.UI{Container: root}
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
