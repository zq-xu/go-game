package setting

import (
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/internal/shooter/settings"
	"github.com/zq-xu/go-game/internal/shooter/stages"
	"github.com/zq-xu/go-game/internal/shooter/ui/components"
	"github.com/zq-xu/go-game/internal/shooter/ui/layout"
)

type settingStage struct {
	ui *ebitenui.UI

	settingWidth int

	stages.BaseStage
}

// NewSettingStage
func NewSettingStage(ctx stages.StageContext) *settingStage {
	s := &settingStage{
		settingWidth: settings.GetSettings().ScreenWidth() / 3,
		BaseStage:    *stages.NewBaseStage(ctx),
	}

	s.initUI()
	return s
}

func (g *settingStage) StageName() stages.StageName {
	return stages.SettingStage
}

func (g *settingStage) Update() error {
	g.ui.Update()
	return nil
}

func (g *settingStage) Draw(screen *ebiten.Image) {
	g.ui.Draw(screen)
}

func (g *settingStage) initUI() {
	root := layout.NewCenterRowLayout(g.settingWidth, 10, nil, func(c *widget.Container) {
		c.AddChild(components.NewMenuButton("Back",
			widget.ButtonOpts.ClickedHandler(
				func(args *widget.ButtonClickedEventArgs) {
					g.Context().SetCurrentGameStage(stages.GamingStage)
				},
			)))
	})

	g.ui = &ebitenui.UI{Container: root}
}
