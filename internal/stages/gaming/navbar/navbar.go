package navbar

import (
	"fmt"
	"image/color"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/internal/listener"
	"github.com/zq-xu/go-game/internal/stages"
	"github.com/zq-xu/go-game/internal/ui/components"
	"github.com/zq-xu/go-game/internal/ui/layout"
	"github.com/zq-xu/go-game/pkg/graphics"
)

type Navbar interface {
	Update()
	Draw(screen *ebiten.Image)
}

type navbar struct {
	ctx stages.StageContext

	ui *ebitenui.UI

	shotUFOText   *widget.Text
	settingButton *widget.Button
}

func NewNavbar(ctx stages.StageContext) *navbar {
	g := &navbar{ctx: ctx}

	g.ui = &ebitenui.UI{
		Container: layout.NewRightTopRowLayout(50, 50, g.Add),
	}
	return g
}

func (g *navbar) Update() {
	g.ui.Update()
}

func (g *navbar) Draw(screen *ebiten.Image) {
	g.shotUFOText.Label = fmt.Sprintf("Shot: %d", listener.GetListener().GameDataListener().GetShotUFO())
	g.ui.Draw(screen)
}

func (g *navbar) Add(c *widget.Container) {
	c.AddChild(g.newShotUFOText())
	c.AddChild(g.newSettingButton())
}

func (g *navbar) newShotUFOText() *widget.Text {
	g.shotUFOText = components.NewCenterText(
		fmt.Sprintf("Shot: %d", listener.GetListener().GameDataListener().GetShotUFO()),
		graphics.GetFont(),
		color.Black,
	)
	return g.shotUFOText
}

func (g *navbar) newSettingButton() *widget.Button {
	g.settingButton = components.NewSettingButton(
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			g.ctx.SetCurrentGameStage(stages.SettingStage)
		}))
	return g.settingButton
}
