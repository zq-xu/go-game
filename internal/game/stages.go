package game

import (
	"github.com/zq-xu/go-game/internal/stages"
	"github.com/zq-xu/go-game/internal/stages/beginning"
	"github.com/zq-xu/go-game/internal/stages/ending"
	"github.com/zq-xu/go-game/internal/stages/gaming"
	"github.com/zq-xu/go-game/internal/stages/menu"
	"github.com/zq-xu/go-game/internal/stages/pause"
	"github.com/zq-xu/go-game/internal/stages/setting"
)

type stageController struct {
	ctx stages.StageContext

	stageSettings map[stages.StageName]stages.GameStage

	// The current game stage
	stages.GameStage
}

// NewStageController
func NewStageController() *stageController {
	s := &stageController{}

	s.ctx = stages.NewStageContext()
	s.ctx.SetStageReseter(s)

	s.stageSettings = make(map[stages.StageName]stages.GameStage)

	s.appendGameStage(beginning.NewBeginningStage(s.ctx))
	s.appendGameStage(ending.NewEndingStage(s.ctx))
	s.appendGameStage(gaming.NewGamingStage(s.ctx))
	s.appendGameStage(pause.NewPauseStage(s.ctx))
	s.appendGameStage(menu.NewMenuStage(s.ctx))
	s.appendGameStage(setting.NewSettingStage(s.ctx))

	s.GameStage = s.stageSettings[s.ctx.CurrentGameStage()]
	return s
}

func (g *stageController) Reset() {
	for _, v := range g.stageSettings {
		v.Reset()
	}
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
