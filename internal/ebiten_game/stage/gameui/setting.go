package gameui

import (
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/2d-game/internal/ebiten_game/game"
)

type SettingOpt func(s *Setting)

type Setting struct {
	ctx *game.Context

	ui *ebitenui.UI

	settingWidth int

	backButtonOpts []widget.ButtonOpt
}

func WithSettingBackButonOpts(opts ...widget.ButtonOpt) SettingOpt {
	return func(s *Setting) { s.backButtonOpts = opts }
}

func NewSetting(ctx *game.Context, opts ...SettingOpt) *Setting {
	g := &Setting{
		ctx:          ctx,
		settingWidth: ctx.Resource.ScreenWidth / 3,
	}

	for _, fn := range opts {
		fn(g)
	}

	g.ui = g.newUI()
	return g
}

func (g *Setting) Update() error {
	g.ui.Update()
	return nil
}

func (g *Setting) Draw(screen *ebiten.Image) {
	g.ui.Draw(screen)
}

func (g *Setting) newUI() *ebitenui.UI {
	root := g.ctx.Resource.LayoutResource.NewCenterRowLayout(g.settingWidth, 10, nil, func(c *widget.Container) {
		c.AddChild(g.ctx.Resource.ButtonResource.NewMenuButton("Back", g.backButtonOpts...))
	})

	return &ebitenui.UI{Container: root}
}
