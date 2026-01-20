package dungeon

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rotisserie/eris"

	"github.com/zq-xu/go-game/internal/dungeon/stages"
	"github.com/zq-xu/go-game/pkg/stage"
)

// NewGame
func NewGame() (ebiten.Game, error) {
	gameStage, err := stages.NewGameStage()
	if err != nil {
		return nil, eris.Wrap(err, "failed to new game stage")
	}
	sc := stage.NewStageController(gameStage)

	// TODO Add stages
	// sc.AppendStage(gameStage)

	return sc, nil
}
