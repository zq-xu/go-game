package pause

import (
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/internal/shooter/stages"
	"github.com/zq-xu/go-game/internal/shooter/ui/components"
	"github.com/zq-xu/go-game/internal/shooter/ui/layout"
	"github.com/zq-xu/go-game/pkg/event"
	"github.com/zq-xu/go-game/pkg/graphics"
)

const (
	pauseTitle = "Pause"
)

var (
	pauseText = []string{"", "", "PRESS SPACE KEY TO START", "", "", "", ""}
)

type pauseStage struct {
	ui *ebitenui.UI

	shadowDrawer func(screen *ebiten.Image)

	inputListener event.InputListener

	stages.BaseStage
}

// NewPauseStage
func NewPauseStage(ctx stages.StageContext) *pauseStage {
	s := &pauseStage{BaseStage: *stages.NewBaseStage(ctx)}

	s.ui = newPauseUI()
	s.shadowDrawer = components.GenerateShadowDrawerFn(0)

	s.inputListener = event.NewInputListener(func() bool {
		if s.IsStable() && ebiten.IsKeyPressed(ebiten.KeySpace) {
			s.Context().SetCurrentGameStage(stages.GamingStage)
			return true
		}

		return false
	})

	return s
}

func (g *pauseStage) StageName() stages.StageName {
	return stages.PauseStage
}

func (g *pauseStage) Update() error {
	g.ui.Update()
	g.inputListener.Update()
	return nil
}

func (g *pauseStage) Draw(screen *ebiten.Image) {
	g.Context().TempDrawer().Draw(screen)
	g.shadowDrawer(screen)
	g.ui.Draw(screen)
}

func newPauseUI() *ebitenui.UI {
	root := layout.NewCenterRowLayout(400, 10, nil,
		func(c *widget.Container) {
			c.AddChild(components.NewCenterText(pauseTitle, graphics.GetFont(), graphics.GetColor().TextIdleColor()))

			for _, v := range pauseText {
				c.AddChild(components.NewCenterText(v, graphics.GetFont(), graphics.GetColor().TextIdleColor()))
			}
		})

	return &ebitenui.UI{Container: root}
}
