package setting

import (
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zq-xu/go-game/internal/settings"
	"github.com/zq-xu/go-game/internal/ui/components"
	"github.com/zq-xu/go-game/internal/ui/layout"
)

type SettingOpt func(s *Setting)

type Setting struct {
	ui *ebitenui.UI

	settingWidth int

	backButtonOpts []widget.ButtonOpt
}

func WithSettingBackButonOpts(opts ...widget.ButtonOpt) SettingOpt {
	return func(s *Setting) { s.backButtonOpts = opts }
}

func NewSetting(opts ...SettingOpt) *Setting {
	g := &Setting{
		settingWidth: settings.GetSettings().ScreenWidth() / 3,
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
	root := layout.NewCenterRowLayout(g.settingWidth, 10, nil, func(c *widget.Container) {
		c.AddChild(components.NewMenuButton("Back", g.backButtonOpts...))
	})

	return &ebitenui.UI{Container: root}
}
