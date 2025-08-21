package game

import (
	"github.com/zq-xu/go-game/internal/stages"
	_ "github.com/zq-xu/go-game/internal/stages/beginning"
	_ "github.com/zq-xu/go-game/internal/stages/ending"
	_ "github.com/zq-xu/go-game/internal/stages/gaming"
	_ "github.com/zq-xu/go-game/internal/stages/menu"
	_ "github.com/zq-xu/go-game/internal/stages/pause"
	_ "github.com/zq-xu/go-game/internal/stages/setting"
)

type stageController struct {
	stages.GameStage
}

// NewStageController
func NewStageController() *stageController {
	return &stageController{GameStage: stages.GetGameStage(stages.MenuStage)}
}

func (g *stageController) Update() error {
	g.GameStage = g.GameStage.NextGameStage()
	return g.GameStage.Update()
}
