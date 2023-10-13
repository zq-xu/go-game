package stages

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/internal/ebiten_game/game"
)

type Interface interface {
	ebiten.Game

	GoNextStatus() (bool, Interface)
}

type Status struct {
	status game.Status

	preStage  Interface
	nextStage Interface
}

func (s *Status) ReloadStatus() {
	s.status = game.RunningStatus
	s.preStage = nil
	s.nextStage = nil
}

func (s *Status) checkoutPreStage() {
	s.status = game.BackStatus
	s.nextStage = s.preStage
}

func (s *Status) checkoutNextStage(nextStage Interface) {
	s.status = game.NextStatus
	s.nextStage = nextStage
}

func (s *Status) GoNextStatus() (bool, Interface) {
	if s.status == game.RunningStatus {
		return false, nil
	}

	defer s.ReloadStatus()
	return true, s.nextStage
}
