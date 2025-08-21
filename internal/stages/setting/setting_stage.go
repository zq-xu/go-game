package setting

import (
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/internal/stages"
)

type SettingStage struct {
	ui *Setting

	stages.BaseStage
}

func init() {
	stages.Register(stages.SettingStage, NewSettingStage())
}

func NewSettingStage() *SettingStage {
	s := &SettingStage{BaseStage: *stages.NewBaseStage()}
	s.SetCurrentGameStage(s)

	s.ui = NewSetting(
		WithSettingBackButonOpts(
			widget.ButtonOpts.ClickedHandler(
				func(args *widget.ButtonClickedEventArgs) {
					s.SetNexttGameStage(stages.GetGameStage(stages.GamingStage))
				},
			),
		),
	)

	return s
}

func (g *SettingStage) Draw(screen *ebiten.Image) { g.ui.Draw(screen) }
