package gameui

import (
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/2d-game/internal/ebiten_game/game"
)

type NavbarOpt func(nb *Navbar)

type Navbar struct {
	ctx *game.Context

	ui *ebitenui.UI

	settingButtonOpts []widget.ButtonOpt
}

func WithNavbarSettingButonOpts(opts ...widget.ButtonOpt) NavbarOpt {
	return func(nb *Navbar) { nb.settingButtonOpts = opts }
}

func NewNavbar(ctx *game.Context, opts ...NavbarOpt) *Navbar {
	g := &Navbar{ctx: ctx}

	for _, fn := range opts {
		fn(g)
	}

	g.ui = g.newUI()
	return g
}

func (g *Navbar) Update() error {
	g.ui.Update()
	return nil
}

func (g *Navbar) Draw(screen *ebiten.Image) {
	g.ui.Draw(screen)
}

func (g *Navbar) newUI() *ebitenui.UI {
	rootContainer := g.ctx.Resource.LayoutResource.NewRightTopRowLayout(0, 0, func(c *widget.Container) {
		c.AddChild(g.newSettingButton())
	})

	return &ebitenui.UI{Container: rootContainer}
}

func (g *Navbar) newSettingButton() *widget.Button {
	return g.ctx.Resource.ButtonResource.NewSettingButton(g.settingButtonOpts...)
}
