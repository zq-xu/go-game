package menu

import (
	"fmt"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/2d-game/internal/ebiten_game/game"
)

const (
	RunningMenuStatus MenuStatus = iota
	StartMenuStatus
	BackMenuStatus
)

type MenuStatus int

type MainMenu struct {
	ctx *game.Context

	ui *ebitenui.UI

	Status MenuStatus
}

func NewMainMenu(ctx *game.Context) *MainMenu {
	g := &MainMenu{ctx: ctx}
	g.ui = g.mainMenuUI()
	return g
}

func (g *MainMenu) Update() error {
	g.ui.Update()
	return nil
}

func (g *MainMenu) Draw(screen *ebiten.Image) {
	g.ui.Draw(screen)
}

func (g *MainMenu) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.ctx.Resource.Cfg.Layout(outsideWidth, outsideHeight)
}

func (g *MainMenu) mainMenuUI() *ebitenui.UI {
	root := g.ctx.Resource.LayoutResource.NewCenterLayout(400, 10, nil, func(c *widget.Container) {
		c.AddChild(g.newStartButton())

		c.AddChild(g.ctx.Resource.SeparatorResource.NewSeparator(widget.RowLayoutData{
			Stretch: true,
		}))

		s := g.ctx.Resource.ButtonResource.NewButton("Toggle Button",
			widget.ButtonOpts.CursorEnteredHandler(func(args *widget.ButtonHoverEventArgs) { fmt.Println("Cursor Entered: " + args.Button.Text().Label) }),
			widget.ButtonOpts.CursorExitedHandler(func(args *widget.ButtonHoverEventArgs) { fmt.Println("Cursor Exited: " + args.Button.Text().Label) }),
			widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
				g.Status = BackMenuStatus
			}))
		c.AddChild(s)

	})

	return &ebitenui.UI{Container: root}
}

func (g *MainMenu) newStartButton() *widget.Button {
	return g.ctx.Resource.ButtonResource.NewButton("Button Start",
		widget.ButtonOpts.CursorEnteredHandler(func(args *widget.ButtonHoverEventArgs) { fmt.Println("Cursor Entered: " + args.Button.Text().Label) }),
		widget.ButtonOpts.CursorExitedHandler(func(args *widget.ButtonHoverEventArgs) { fmt.Println("Cursor Exited: " + args.Button.Text().Label) }),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			g.Status = StartMenuStatus
		}))
}
