package pause

import (
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/internal/stages"
	"github.com/zq-xu/go-game/internal/ui/components"
	"github.com/zq-xu/go-game/internal/ui/layout"
	"github.com/zq-xu/go-game/pkg/event"
	"github.com/zq-xu/go-game/pkg/graphics"
)

const (
	pauseTitle = "Pause"
)

var (
	pauseText = []string{"", "", "PRESS SPACE KEY TO START", "", "", "", ""}
)

type PauseStage struct {
	ui *ebitenui.UI

	shadowDrawer func(screen *ebiten.Image)

	inputListener event.InputListener

	stages.BaseStage
}

func init() {
	stages.Register(stages.PauseStage, NewPauseStage())
}

func NewPauseStage() *PauseStage {
	s := &PauseStage{BaseStage: *stages.NewBaseStage()}
	s.SetCurrentGameStage(s)

	s.ui = newPauseUI()
	s.shadowDrawer = components.GenerateShadowDrawerFn(0)

	s.inputListener = event.NewInputListener(func() bool {
		if s.IsStable() && ebiten.IsKeyPressed(ebiten.KeySpace) {
			s.SetNexttGameStage(stages.GetGameStage(stages.GamingStage))
			return true
		}

		return false
	})

	return s
}

func (g *PauseStage) Update() error {
	g.ui.Update()
	g.inputListener.Update()
	return nil
}

func (g *PauseStage) Draw(screen *ebiten.Image) {
	stages.GetGameStage(stages.GamingStage).Draw(screen)
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

func (g *PauseStage) ReShown() { g.inputListener.Reload() }
