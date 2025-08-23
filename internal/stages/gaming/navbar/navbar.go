package navbar

import (
	"fmt"
	"image/color"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/internal/data"
	"github.com/zq-xu/go-game/internal/stages"
	"github.com/zq-xu/go-game/internal/ui/components"
	"github.com/zq-xu/go-game/internal/ui/layout"
	"github.com/zq-xu/go-game/pkg/graphics"
)

const (
	fontSize = 20
)

type Navbar interface {
	Update()
	Draw(screen *ebiten.Image)
}

type navbar struct {
	ctx      stages.StageContext
	gameData data.Data

	ui *ebitenui.UI

	shotUFOText   *widget.Text
	settingButton *widget.Button
}

func NewNavbar(ctx stages.StageContext, gameData data.Data) *navbar {
	g := &navbar{
		ctx:      ctx,
		gameData: gameData,
	}

	g.ui = &ebitenui.UI{Container: layout.NewRightTopRowLayout(50, 50, g.Add)}
	return g
}

func (g *navbar) Update() {
	g.ui.Update()
}

func (g *navbar) Draw(screen *ebiten.Image) {
	g.shotUFOText.Label = g.generateShotLabel()
	g.ui.Draw(screen)
}

func (g *navbar) Add(c *widget.Container) {
	c.AddChild(g.newShotUFOText())
	c.AddChild(g.newSettingButton())
}

func (g *navbar) newShotUFOText() *widget.Text {
	g.shotUFOText = components.NewCenterText(
		g.generateShotLabel(),
		graphics.NewNotoSansFont(fontSize),
		color.White,
	)
	return g.shotUFOText
}

func (g *navbar) generateShotLabel() string {
	return fmt.Sprintf("Shot: %d", g.gameData.GetShotUFO())
}

func (g *navbar) newSettingButton() *widget.Button {
	g.settingButton = components.NewSettingButton(
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			g.ctx.SetCurrentGameStage(stages.SettingStage)
		}))
	return g.settingButton
}
