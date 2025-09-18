package ending

import (
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/internal/shooter/stages"
	"github.com/zq-xu/go-game/internal/shooter/status"
	"github.com/zq-xu/go-game/internal/shooter/ui/components"
	"github.com/zq-xu/go-game/internal/shooter/ui/layout"
	"github.com/zq-xu/go-game/pkg/config"
	"github.com/zq-xu/go-game/pkg/graphics"
)

var (
	SuccessText = "SUCCESS"
	FailureText = "FAILED"

	resultTextSet = map[status.Status]string{
		status.SuccessStatus: SuccessText,
		status.FailStatus:    FailureText,
	}

	titleTexts = []string{"GAME OVER"}
)

type endingStage struct {
	ui *ebitenui.UI

	statusText *widget.Text
	stages.BaseStage
}

func NewEndingStage(ctx stages.StageContext) *endingStage {
	s := &endingStage{BaseStage: *stages.NewBaseStage(ctx)}
	s.initUi()
	return s
}

func (g *endingStage) StageName() stages.StageName {
	return stages.EndingStage
}

func (g *endingStage) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		g.Context().Reset()
		g.Context().SetCurrentGameStage(stages.GamingStage)
	}

	g.ui.Update()
	return nil
}

func (g *endingStage) Draw(screen *ebiten.Image) {
	g.statusText.Label = resultTextSet[g.Context().Status()]
	g.ui.Draw(screen)
}

func (g *endingStage) initUi() {
	g.statusText = components.NewCenterText("", graphics.GetFont(), graphics.GetColor().TextIdleColor())

	root := layout.NewCenterRowLayout(400, 10, nil, func(c *widget.Container) {
		for _, v := range titleTexts {
			c.AddChild(components.NewCenterText(v, graphics.GetFont(), graphics.GetColor().TextIdleColor()))
		}

		c.AddChild(g.statusText)

		for _, v := range config.Cfg.StartHintTexts {
			c.AddChild(components.NewCenterText(v, graphics.GetFont(), graphics.GetColor().TextIdleColor()))
		}
	})

	g.ui = &ebitenui.UI{Container: root}
}
