package gameui

import (
	"fmt"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/2d-game/internal/ebiten_game/game"
)

type Navbar struct {
	ctx *game.Context

	ui *ebitenui.UI
}

func NewNavbar(ctx *game.Context) *Navbar {
	g := &Navbar{
		ctx: ctx,
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
		c.AddChild(g.ctx.Resource.ButtonResource.NewSettingButton(
			widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
				fmt.Print("click setting")
			})))
	})

	return &ebitenui.UI{Container: rootContainer}
}
