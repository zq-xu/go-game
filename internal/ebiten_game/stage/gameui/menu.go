package gameui

import (
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/2d-game/internal/ebiten_game/game"
)

type MenuOpt func(m *Menu)

type Menu struct {
	ctx *game.Context

	ui *ebitenui.UI

	menuWidth int

	nextButtonOpts []widget.ButtonOpt

	Status game.Status
}

func WithMenuNextButonOpts(opts ...widget.ButtonOpt) MenuOpt {
	return func(m *Menu) { m.nextButtonOpts = opts }
}

func NewMenu(ctx *game.Context, opts ...MenuOpt) *Menu {
	g := &Menu{
		ctx:       ctx,
		menuWidth: ctx.Resource.ScreenWidth / 3,
	}

	for _, fn := range opts {
		fn(g)
	}

	g.ui = g.newUI()
	return g
}

func (g *Menu) Update() error {
	g.ui.Update()
	return nil
}

func (g *Menu) Draw(screen *ebiten.Image) {
	g.ui.Draw(screen)
}

func (g *Menu) newUI() *ebitenui.UI {
	root := g.ctx.Resource.LayoutResource.NewCenterRowLayout(g.menuWidth, 10, nil, func(c *widget.Container) {
		c.AddChild(g.ctx.Resource.ButtonResource.NewMenuButton("Start"))
		c.AddChild(g.ctx.Resource.SeparatorResource.NewSeparator(widget.RowLayoutData{Stretch: true}))
		c.AddChild(g.ctx.Resource.ButtonResource.NewMenuButton("Next", g.nextButtonOpts...))
		c.AddChild(g.ctx.Resource.SeparatorResource.NewSeparator(widget.RowLayoutData{Stretch: true}))
		c.AddChild(g.ctx.Resource.ButtonResource.NewMenuButton("Back"))
	})

	return &ebitenui.UI{Container: root}
}
