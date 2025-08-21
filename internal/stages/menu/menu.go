package menu

import (
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/internal/settings"
	"github.com/zq-xu/go-game/internal/stages"
	"github.com/zq-xu/go-game/internal/ui/components"
	"github.com/zq-xu/go-game/internal/ui/layout"
)

type MenuStage struct {
	ui *ebitenui.UI
	stages.BaseStage
}

func init() {
	stages.Register(stages.MenuStage, NewMenuStage())
}

// NewMenuStage
func NewMenuStage() *MenuStage {
	s := &MenuStage{BaseStage: *stages.NewBaseStage()}
	s.SetCurrentGameStage(s)

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

func (g *MenuStage) startButton() *widget.Button {
	return components.NewMenuButton("Start",
		widget.ButtonOpts.ClickedHandler(
			func(args *widget.ButtonClickedEventArgs) {
				g.SetNexttGameStage(stages.GetGameStage(stages.BeginningStage))
			}))

}

func (g *MenuStage) nextButton() *widget.Button {
	return components.NewMenuButton("Next",
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			g.SetNexttGameStage(stages.GetGameStage(stages.GamingStage))
		}))

}

func (g *MenuStage) backButton() *widget.Button {
	return components.NewMenuButton("Back")
}

func (g *MenuStage) Update() error {
	g.ui.Update()
	return nil
}

func (g *MenuStage) Draw(screen *ebiten.Image) { g.ui.Draw(screen) }
