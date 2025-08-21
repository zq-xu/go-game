package beginning

import (
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/internal/config"
	"github.com/zq-xu/go-game/internal/settings"
	"github.com/zq-xu/go-game/internal/stages"
	"github.com/zq-xu/go-game/internal/ui/components"
	"github.com/zq-xu/go-game/internal/ui/layout"
	"github.com/zq-xu/go-game/pkg/graphics"
)

type beginningStage struct {
	ui *ebitenui.UI
	stages.BaseStage
}

func init() {
	stages.Register(stages.BeginningStage, NewBeginningStage())
}

func NewBeginningStage() *beginningStage {
	s := &beginningStage{
		ui:        NewBeginningUI(),
		BaseStage: *stages.NewBaseStage(),
	}
	s.SetCurrentGameStage(s)
	return s
}

func (g *beginningStage) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		g.SetNexttGameStage(stages.GetGameStage(stages.GamingStage))
	}

	g.ui.Update()
	return nil
}

func (g *beginningStage) Draw(screen *ebiten.Image) {
	g.ui.Draw(screen)
}

func (g *beginningStage) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return settings.GetSettings().Layout(outsideWidth, outsideHeight)
}

func NewBeginningUI() *ebitenui.UI {
	root := layout.NewCenterRowLayout(400, 10, nil, func(c *widget.Container) {
		c.AddChild(components.NewCenterText(config.Cfg.Title, graphics.GetFont(), graphics.GetColor().TextIdleColor()))
		c.AddChild(components.NewCenterText(config.Cfg.AuthorText, graphics.GetFont(), graphics.GetColor().TextIdleColor()))

		for _, v := range config.Cfg.StartHintTexts {
			c.AddChild(components.NewCenterText(v, graphics.GetFont(), graphics.GetColor().TextIdleColor()))
		}
	})

	return &ebitenui.UI{Container: root}
}
