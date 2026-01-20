package npc

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rotisserie/eris"

	"github.com/zq-xu/go-game/internal/dungeon/core/actor"
	"github.com/zq-xu/go-game/internal/dungeon/core/controls"
	"github.com/zq-xu/go-game/internal/dungeon/core/dialog"
	"github.com/zq-xu/gotools/logx"
)

type NPC interface {
	actor.Actor

	controls.ApprochObject
}

type npc struct {
	cfg *NPCConfig

	actor.Actor

	approaching bool
	dialog.Dialog
}

// LoadNPCs
func LoadNPCs() ([]actor.Actor, error) {
	list := make([]actor.Actor, 0)

	for k, v := range NPCConfigSet {
		logx.Logger.Infof("Loading NPC %s", k)
		npc, err := NewNPC(v)
		if err != nil {
			return nil, eris.Wrapf(err, "failed to new %s", v.Name)
		}

		list = append(list, npc)
		logx.Logger.Infof("Loaded NPC %s", k)

		controls.RegisterNPC(npc.Name(), npc)
	}

	return list, nil
}

// NewNPC
func NewNPC(cfg *NPCConfig) (NPC, error) {
	a, err := actor.NewActor(&cfg.Option)
	if err != nil {
		return nil, eris.Wrap(err, "failed to new actor")
	}

	d, err := dialog.NewDialog()
	if err != nil {
		return nil, eris.Wrap(err, "failed to new dialog")
	}

	mg := &npc{
		cfg:    cfg,
		Actor:  a,
		Dialog: d,
	}
	return mg, nil
}

func (s *npc) Update() {
	s.Idle()
}

func (s *npc) Draw(screen *ebiten.Image) {
	s.Actor.Draw(screen)
	s.drawDialog(screen)
}

func (s *npc) drawDialog(screen *ebiten.Image) {
	if !s.approaching {
		return
	}

	leftX, leftTopY := s.LeftTop()
	x := leftX + s.cfg.Width + s.cfg.DialogOffsetToRightCenter.X
	y := leftTopY - s.cfg.DialogOffsetToRightCenter.Y

	s.Dialog.DrawIcon(screen, x, y)

}

func (s *npc) MoveNearby() { s.approaching = true }
func (s *npc) MoveAway()   { s.approaching = false }
