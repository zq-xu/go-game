package gaming

import (
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/internal/config"
	"github.com/zq-xu/go-game/internal/stages"
	"github.com/zq-xu/go-game/pkg/event"
)

type GamingStage struct {
	runtime       *Runtime
	navbar        Navbar
	metric        *Metric
	inputListener event.InputListener

	stages.BaseStage
}

func init() {
	stages.Register(stages.GamingStage, NewGamingStage())
}

func NewGamingStage() *GamingStage {
	s := &GamingStage{BaseStage: *stages.NewBaseStage()}
	s.SetCurrentGameStage(s)

	s.runtime = NewRuntime(
		WithRuntimeCollission(func(b bool) {
			if !b {
				return
			}
			s.SetNexttGameStage(stages.GetGameStage(stages.EndingStage))
		}),
	)

	s.navbar = NewNavbar(
		WithNavbarSettingButonOpts(
			widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
				s.SetNexttGameStage(stages.GetGameStage(stages.SettingStage))
			}),
		))

	s.inputListener = event.NewInputListener(func() bool {
		if s.IsStable() && ebiten.IsKeyPressed(ebiten.KeySpace) {
			s.SetNexttGameStage(stages.GetGameStage(stages.PauseStage))
			return true
		}

		return false
	})

	s.metric = NewMetric()

	return s
}

func (g *GamingStage) Update() error {
	err := g.runtime.Update()
	if err != nil {
		return err
	}

	err = g.navbar.Update()
	if err != nil {
		return err
	}

	g.inputListener.Update()
	return nil
}

func (g *GamingStage) Draw(screen *ebiten.Image) {
	screen.Fill(config.Cfg.BgColor)

	g.runtime.Draw(screen)
	g.navbar.Draw(screen)
	g.metric.Draw(screen)
}

func (g *GamingStage) ReShown() { g.inputListener.Reload() }
