package menu

import (
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/internal/shooter/settings"
	"github.com/zq-xu/go-game/internal/shooter/stages"
	"github.com/zq-xu/go-game/internal/shooter/ui/components"
	"github.com/zq-xu/go-game/internal/shooter/ui/layout"
)

type menuStage struct {
	ui *ebitenui.UI
	stages.BaseStage
}

// NewMenuStage
func NewMenuStage(ctx stages.StageContext) *menuStage {
	s := &menuStage{BaseStage: *stages.NewBaseStage(ctx)}

	menuWidth := settings.GetSettings().ScreenWidth() / 3
	root := layout.NewCenterRowLayout(menuWidth, 10, nil, func(c *widget.Container) {
		c.AddChild(s.startButton())
		c.AddChild(components.NewSeparator(widget.RowLayoutData{Stretch: true}))
		c.AddChild(s.nextButton())
		c.AddChild(components.NewSeparator(widget.RowLayoutData{Stretch: true}))
		c.AddChild(s.backButton())
	})

	s.ui = &ebitenui.UI{Container: root}
	return s
}

func (g *menuStage) StageName() stages.StageName {
	return stages.MenuStage
}

func (g *menuStage) Update() error {
	g.ui.Update()
	return nil
}

func (g *menuStage) Draw(screen *ebiten.Image) {
	g.ui.Draw(screen)
}

func (g *menuStage) startButton() *widget.Button {
	return components.NewMenuButton("Start",
		widget.ButtonOpts.ClickedHandler(
			func(args *widget.ButtonClickedEventArgs) {
				g.Context().SetCurrentGameStage(stages.BeginningStage)
			}))

}

func (g *menuStage) nextButton() *widget.Button {
	return components.NewMenuButton("Next",
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			g.Context().SetCurrentGameStage(stages.GamingStage)
		}))

}

func (g *menuStage) backButton() *widget.Button {
	return components.NewMenuButton("Back")
}
