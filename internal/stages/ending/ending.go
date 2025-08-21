package ending

import (
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/internal/config"
	"github.com/zq-xu/go-game/internal/settings"
	"github.com/zq-xu/go-game/internal/stages"
	"github.com/zq-xu/go-game/internal/status"
	"github.com/zq-xu/go-game/internal/ui/components"
	"github.com/zq-xu/go-game/internal/ui/layout"
	"github.com/zq-xu/go-game/pkg/graphics"
)

var (
	SuccessText = "SUCCESS"
	FailureText = "FAILED"

	resultTextSet = map[status.Status]string{
		status.SuccessStatus: SuccessText,
		status.FailStatus:    FailureText,
	}
)

type EndingStage struct {
	ui *ebitenui.UI

	stages.BaseStage
}

func init() {
	stages.Register(stages.EndingStage, NewEndingStage())
}

func NewEndingStage() *EndingStage {
	s := &EndingStage{
		ui:        NewEndingUI(),
		BaseStage: *stages.NewBaseStage(),
	}
	s.SetCurrentGameStage(s)
	return s
}

func (g *EndingStage) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		g.SetNexttGameStage(stages.GetGameStage(stages.GamingStage))
	}

	g.ui.Update()

	return nil
}

func (g *EndingStage) Draw(screen *ebiten.Image) {
	g.ui.Draw(screen)
}

func (g *EndingStage) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return settings.GetSettings().Layout(outsideWidth, outsideHeight)
}

func NewEndingUI() *ebitenui.UI {
	titleTexts := []string{"GAME OVER", resultTextSet[status.SuccessStatus]}

	root := layout.NewCenterRowLayout(400, 10, nil, func(c *widget.Container) {
		for _, v := range titleTexts {
			c.AddChild(components.NewCenterText(v, graphics.GetFont(), graphics.GetColor().TextIdleColor()))
		}

		for _, v := range config.Cfg.StartHintTexts {
			c.AddChild(components.NewCenterText(v, graphics.GetFont(), graphics.GetColor().TextIdleColor()))
		}
	})

	return &ebitenui.UI{Container: root}
}
