package gameui

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

	menuWith int

	Status MenuStatus
}

func NewMainMenu(ctx *game.Context) *MainMenu {
	g := &MainMenu{ctx: ctx}

	g.menuWith = ctx.Resource.ScreenWidth / 3
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
	return g.ctx.Resource.Layout(outsideWidth, outsideHeight)
}

func (g *MainMenu) mainMenuUI() *ebitenui.UI {
	root := g.ctx.Resource.LayoutResource.NewCenterRowLayout(g.menuWith, 10, nil, func(c *widget.Container) {
		c.AddChild(g.ctx.Resource.ButtonResource.NewMenuButton("Button Start",
			widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
				g.Status = StartMenuStatus
			})))

		c.AddChild(g.ctx.Resource.SeparatorResource.NewSeparator(widget.RowLayoutData{
			Stretch: true,
		}))

		s := g.ctx.Resource.ButtonResource.NewMenuButton("Toggle Button",
			widget.ButtonOpts.CursorEnteredHandler(func(args *widget.ButtonHoverEventArgs) { fmt.Println("Cursor Entered: " + args.Button.Text().Label) }),
			widget.ButtonOpts.CursorExitedHandler(func(args *widget.ButtonHoverEventArgs) { fmt.Println("Cursor Exited: " + args.Button.Text().Label) }),
			widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
				g.Status = BackMenuStatus
			}))
		c.AddChild(s)

	})

	return &ebitenui.UI{Container: root}
}
