package setting

import (
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/internal/stages"
)

type settingStage struct {
	ui *Setting

	stages.BaseStage
}

// NewSettingStage
func NewSettingStage(ctx stages.StageContext) *settingStage {
	s := &settingStage{BaseStage: *stages.NewBaseStage(ctx)}

	s.ui = NewSetting(
		WithSettingBackButonOpts(
			widget.ButtonOpts.ClickedHandler(
				func(args *widget.ButtonClickedEventArgs) {
					s.Context().SetCurrentGameStage(stages.GamingStage)
				},
			),
		),
	)

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
