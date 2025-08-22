package beginning

import (
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/internal/config"
	"github.com/zq-xu/go-game/internal/stages"
	"github.com/zq-xu/go-game/internal/ui/components"
	"github.com/zq-xu/go-game/internal/ui/layout"
	"github.com/zq-xu/go-game/pkg/graphics"
)

type beginningStage struct {
	ui *ebitenui.UI
	stages.BaseStage
}

func NewBeginningStage(ctx stages.StageContext) stages.GameStage {
	s := &beginningStage{
		ui:        newBeginningUI(),
		BaseStage: *stages.NewBaseStage(ctx),
	}

	return s
}

func (g *beginningStage) StageName() stages.StageName {
	return stages.BeginningStage
}

func (g *beginningStage) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		g.Context().SetCurrentGameStage(stages.GamingStage)
	}

	g.ui.Update()
	return nil
}

func (g *beginningStage) Draw(screen *ebiten.Image) {
	g.ui.Draw(screen)
}

func newBeginningUI() *ebitenui.UI {
	root := layout.NewCenterRowLayout(400, 10, nil, func(c *widget.Container) {
		c.AddChild(components.NewCenterText(config.Cfg.Title, graphics.GetFont(), graphics.GetColor().TextIdleColor()))
		c.AddChild(components.NewCenterText(config.Cfg.AuthorText, graphics.GetFont(), graphics.GetColor().TextIdleColor()))

		for _, v := range config.Cfg.StartHintTexts {
			c.AddChild(components.NewCenterText(v, graphics.GetFont(), graphics.GetColor().TextIdleColor()))
		}
	})

	return &ebitenui.UI{Container: root}
}
