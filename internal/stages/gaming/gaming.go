package gaming

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/internal/stages"
	"github.com/zq-xu/go-game/internal/stages/gaming/gamerun"
	"github.com/zq-xu/go-game/internal/stages/gaming/navbar"
	"github.com/zq-xu/go-game/pkg/event"
)

type gamingStage struct {
	gamerun       gamerun.GameRun
	navbar        navbar.Navbar
	metric        *Metric
	inputListener event.InputListener

	stages.BaseStage
}

func NewGamingStage(ctx stages.StageContext) *gamingStage {
	s := &gamingStage{BaseStage: *stages.NewBaseStage(ctx)}

	s.initGameRun()
	s.initNavbar()
	s.initInputListener()

	s.metric = NewMetric()
	return s
}

func (g *gamingStage) StageName() stages.StageName {
	return stages.GamingStage
}

func (g *gamingStage) Update() error {
	g.gamerun.Update()
	g.navbar.Update()
	g.inputListener.Update()
	return nil
}

func (g *gamingStage) Draw(screen *ebiten.Image) {
	g.gamerun.Draw(screen)
	g.navbar.Draw(screen)
	g.metric.Draw(screen)
}

func (g *gamingStage) Reset() {
	g.initGameRun()
	g.BaseStage.Reset()
}

func (g *gamingStage) initGameRun() {
	g.gamerun = gamerun.NewGameRun(g.Context())
}

func (g *gamingStage) initNavbar() {
	g.navbar = navbar.NewNavbar(g.Context())
}

func (g *gamingStage) initInputListener() {
	g.inputListener = event.NewInputListener(func() bool {
		if g.IsStable() && ebiten.IsKeyPressed(ebiten.KeySpace) {
			g.Context().SetTempDrawer(g)
			g.Context().SetCurrentGameStage(stages.PauseStage)
			return true
		}

		return false
	})
}
