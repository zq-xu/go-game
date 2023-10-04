package stage

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zq-xu/2d-game/internal/ebiten_game/game"
)

type Interface interface {
	ebiten.Game

	GoNextStatus() (bool, Interface)
}

type StageController struct {
	ctx *game.Context

	stage Interface
}

func NewStageController(ctx *game.Context) *StageController {
	s := &StageController{}

	s.ctx = ctx
	// s.stage = NewBeginningStage(loader)
	s.stage = NewMenuStage(ctx)

	return s
}

func (s *StageController) Update() error {
	err := s.stage.Update()
	if err != nil {
		return err
	}

	b, i := s.stage.GoNextStatus()
	if b {
		s.stage = i
	}

	return nil
}

func (s *StageController) Draw(screen *ebiten.Image) {
	s.stage.Draw(screen)
}

func (s *StageController) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return s.stage.Layout(outsideWidth, outsideHeight)
}
