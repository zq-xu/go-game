package stage

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zq-xu/2d-game/internal/ebiten_game/loader"
)

const (
	RunningStageStatus StageStatus = iota
	SuccessStageStatus
	FailStageStatus
	ReturnBackStageStatus
)

type StageStatus int

type Interface interface {
	ebiten.Game

	GoNextStatus() (bool, Interface)
}

type StageController struct {
	loader *loader.Loader
	stage  Interface
}

func NewStageController(loader *loader.Loader) *StageController {
	return &StageController{
		stage: NewBeginningStage(loader),
	}
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
