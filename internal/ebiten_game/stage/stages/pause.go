package stages

import (
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/2d-game/internal/ebiten_game/game"
	"github.com/zq-xu/2d-game/internal/ebiten_game/stage/listener"
)

const (
	pauseTitle = "Pause"
)

var (
	pauseText = []string{"", "", "PRESS SPACE KEY TO START", "", "", "", ""}
)

type PauseStage struct {
	ctx *game.Context

	ui *ebitenui.UI

	shadowDrawer func(screen *ebiten.Image)

	inputLitener *listener.InputListener

	Status
}

func NewPauseStage(ctx *game.Context, preStage Interface) *PauseStage {
	g := &PauseStage{
		ctx: ctx,
		ui:  NewPauseUI(ctx),
		Status: Status{
			preStage: preStage,
		},
	}

	g.inputLitener = listener.NewInputListener(ctx, func() bool {
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			g.checkoutPreStage()
			return true
		}

		return false
	})

	g.shadowDrawer = g.ctx.Resource.ShadowResource.GenerateShadowDrawerFn(0)

	return g
}

func (g *PauseStage) Update() error {
	g.ui.Update()

	g.inputLitener.Update()
	return nil
}

func (g *PauseStage) Draw(screen *ebiten.Image) {
	g.preStage.Draw(screen)

	g.shadowDrawer(screen)

	g.ui.Draw(screen)
}

func (g *PauseStage) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.ctx.Resource.Layout(outsideWidth, outsideHeight)
}

func NewPauseUI(ctx *game.Context) *ebitenui.UI {
	root := ctx.Resource.LayoutResource.NewCenterRowLayout(400, 10, nil, func(c *widget.Container) {
		c.AddChild(ctx.Resource.TextResource.NewCenterText(pauseTitle, ctx.Resource.FontLoader.TitleFace()))

		for _, v := range pauseText {
			c.AddChild(ctx.Resource.TextResource.NewCenterText(v, ctx.Resource.FontLoader.Face()))
		}
	})

	return &ebitenui.UI{Container: root}
}
