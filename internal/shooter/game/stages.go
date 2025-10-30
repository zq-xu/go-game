package game

import (
	"github.com/rotisserie/eris"

	"github.com/zq-xu/go-game/internal/shooter/data"
	"github.com/zq-xu/go-game/internal/shooter/stages"
	"github.com/zq-xu/go-game/internal/shooter/stages/beginning"
	"github.com/zq-xu/go-game/internal/shooter/stages/ending"
	"github.com/zq-xu/go-game/internal/shooter/stages/gaming"
	"github.com/zq-xu/go-game/internal/shooter/stages/menu"
	"github.com/zq-xu/go-game/internal/shooter/stages/pause"
	"github.com/zq-xu/go-game/internal/shooter/stages/setting"
)

type stageController struct {
	ctx      stages.StageContext
	gameData data.Data

	stageSettings map[stages.StageName]stages.GameStage

	// The current game stage
	stages.GameStage
}

// NewStageController
func NewStageController(gameData data.Data) (*stageController, error) {
	s := &stageController{
		gameData: gameData,
	}

	s.ctx = stages.NewStageContext()
	s.ctx.SetStageReseter(s)

	s.stageSettings = make(map[stages.StageName]stages.GameStage)
	s.appendGameStage(beginning.NewBeginningStage(s.ctx))
	s.appendGameStage(ending.NewEndingStage(s.ctx))

	gamingStage, err := gaming.NewGamingStage(s.ctx, gameData)
	if err != nil {
		return nil, eris.Wrap(err, "init game run failed.")
	}
	s.appendGameStage(gamingStage)

	s.appendGameStage(pause.NewPauseStage(s.ctx))
	s.appendGameStage(menu.NewMenuStage(s.ctx))
	s.appendGameStage(setting.NewSettingStage(s.ctx))

	s.GameStage = s.stageSettings[s.ctx.CurrentGameStage()]
	return s, nil
}

func (g *stageController) Reset() error {
	g.gameData.Reset()

	for _, v := range g.stageSettings {
		err := v.Reset()
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *stageController) Update() error {
	g.refreshGameStage()
	return g.GameStage.Update()
}

func (g *stageController) appendGameStage(s stages.GameStage) {
	g.stageSettings[s.StageName()] = s
}

func (g *stageController) refreshGameStage() {
	if g.GameStage.StageName() == g.ctx.CurrentGameStage() {
		return
	}

	g.GameStage = g.stageSettings[g.ctx.CurrentGameStage()]
	g.GameStage.Reload()
}
