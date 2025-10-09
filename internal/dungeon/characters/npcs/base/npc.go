package base

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rotisserie/eris"

	"github.com/zq-xu/go-game/internal/dungeon/core/actor"
	"github.com/zq-xu/go-game/internal/dungeon/core/dialog"
	"github.com/zq-xu/go-game/internal/dungeon/types"
	"github.com/zq-xu/go-game/pkg/event/input"
)

type npc struct {
	cfg *NPCConfig
	actor.Actor

	approaching bool
	dialog.Dialog
}

type NPCConfig struct {
	Name string

	Width, Height             float64
	StartPoint                *image.Point
	DialogOffsetToRightCenter *types.Offset
	IdleImagePaths            map[input.Direction]string
}

func NewNPC(cfg *NPCConfig) (types.NPC, error) {
	a, err := actor.NewActor(&actor.Option{
		Name:           cfg.Name,
		StartPoint:     cfg.StartPoint,
		Width:          cfg.Width,
		Height:         cfg.Height,
		IdleImagePaths: cfg.IdleImagePaths,
	})
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
