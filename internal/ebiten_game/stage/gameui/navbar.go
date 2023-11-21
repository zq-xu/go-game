package gameui

import (
	"fmt"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/internal/ebiten_game/game"
)

type NavbarOpt func(nb *Navbar)

type Navbar struct {
	ctx *game.Context

	ui *ebitenui.UI

	settingButtonOpts []widget.ButtonOpt

	shotUFOText *widget.Text
}

func WithNavbarSettingButonOpts(opts ...widget.ButtonOpt) NavbarOpt {
	return func(nb *Navbar) { nb.settingButtonOpts = opts }
}

func NewNavbar(ctx *game.Context, opts ...NavbarOpt) *Navbar {
	g := &Navbar{ctx: ctx}

	for _, fn := range opts {
		fn(g)
	}

	g.initializeUI()
	return g
}

func (g *Navbar) Update() error {
	g.ui.Update()
	return nil
}

func (g *Navbar) Draw(screen *ebiten.Image) {
	g.shotUFOText.Label = fmt.Sprintf("Shot: %d", g.ctx.Listener.GameDataListener().GetShotUFO())
	g.ui.Draw(screen)
}

func (g *Navbar) initializeUI() {
	g.shotUFOText = g.ctx.Resource.TextResource.NewCenterText(
		fmt.Sprintf("Shot: %d", g.ctx.Listener.GameDataListener().GetShotUFO()),
		g.ctx.Resource.FontLoader.TitleFace(),
	)

	g.ui = &ebitenui.UI{
		Container: g.ctx.Resource.LayoutResource.NewRightTopRowLayout(100, 0, func(c *widget.Container) {
			c.AddChild(g.shotUFOText)
			c.AddChild(g.newSettingButton())
		}),
	}
}

func (g *Navbar) newSettingButton() *widget.Button {
	return g.ctx.Resource.ButtonResource.NewSettingButton(g.settingButtonOpts...)
}
