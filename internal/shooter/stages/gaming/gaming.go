package gaming

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rotisserie/eris"

	"github.com/zq-xu/go-game/internal/shooter/data"
	"github.com/zq-xu/go-game/internal/shooter/stages"
	"github.com/zq-xu/go-game/internal/shooter/stages/gaming/gamerun"
	"github.com/zq-xu/go-game/internal/shooter/stages/gaming/navbar"
	"github.com/zq-xu/go-game/pkg/event"
)

type gamingStage struct {
	gameData data.Data

	gamerun       gamerun.GameRun
	navbar        navbar.Navbar
	metrics       *metrics
	inputListener event.InputListener

	stages.BaseStage
}

func NewGamingStage(ctx stages.StageContext, gameData data.Data) (*gamingStage, error) {
	s := &gamingStage{
		gameData:  gameData,
		BaseStage: *stages.NewBaseStage(ctx),
	}

	err := s.initGameRun()
	if err != nil {
		return nil, eris.Wrap(err, "init game run failed.")
	}
	s.initNavbar()
	s.initInputListener()

	s.metrics = NewMetric(gameData)
	return s, nil
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
	g.metrics.Draw(screen)
}

func (g *gamingStage) Reset() error {
	err := g.initGameRun()
	if err != nil {
		return err
	}

	return g.BaseStage.Reset()
}

func (g *gamingStage) initGameRun() error {
	var err error
	g.gamerun, err = gamerun.NewGameRun(g.Context(), g.gameData)
	return err
}

func (g *gamingStage) initNavbar() {
	g.navbar = navbar.NewNavbar(g.Context(), g.gameData)
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
