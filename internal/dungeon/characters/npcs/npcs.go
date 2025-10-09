package npcs

import (
	"fmt"
	"image"

	"github.com/rotisserie/eris"

	"github.com/zq-xu/go-game/assets/dungeon"
	"github.com/zq-xu/go-game/internal/dungeon/characters/npcs/base"
	"github.com/zq-xu/go-game/internal/dungeon/core/controls"
	"github.com/zq-xu/go-game/internal/dungeon/types"
	"github.com/zq-xu/go-game/pkg/event/input"
)

func NewGrocery() (types.NPC, error) {
	return base.NewNPC(&base.NPCConfig{
		Name:       "Grocery",
		StartPoint: &image.Point{X: 100, Y: 650},
		Width:      40,
		Height:     45,
		DialogOffsetToRightCenter: &types.Offset{
			X: 0,
			Y: 5,
		},
		IdleImagePaths: map[input.Direction]string{
			input.DefaultDirection: dungeon.GetDungeonImagePath("/characters/npcs/grocery/idle.png"),
		},
	})
}

func NewMoonTower() (types.NPC, error) {
	return base.NewNPC(&base.NPCConfig{
		Name:       "MoonTower",
		StartPoint: &image.Point{X: 550, Y: 50},
		Width:      70,
		Height:     90,
		DialogOffsetToRightCenter: &types.Offset{
			X: 0,
			Y: 5,
		},
		IdleImagePaths: map[input.Direction]string{
			input.DefaultDirection: dungeon.GetDungeonImagePath("/characters/npcs/moon_tower/idle.png"),
		},
	})
}

func NewMageGuardian() (types.NPC, error) {
	return base.NewNPC(&base.NPCConfig{
		Name:       fmt.Sprintf("MageGuardian-%s", "blue"),
		StartPoint: &image.Point{X: 700, Y: 300},
		Width:      35,
		Height:     60,
		DialogOffsetToRightCenter: &types.Offset{
			X: 0,
			Y: 5,
		},
		IdleImagePaths: map[input.Direction]string{
			input.DefaultDirection: dungeon.GetDungeonImagePath("/characters/npcs/mage_guardian/idle.png"),
		},
	})
}

func LoadNPCs() error {
	mg, err := NewMageGuardian()
	if err != nil {
		return eris.Wrap(err, "failed to new mage guardian")
	}
	controls.RegisterNPC(mg)

	mt, err := NewMoonTower()
	if err != nil {
		return eris.Wrap(err, "failed to new moon tower")
	}
	controls.RegisterNPC(mt)

	bs, err := NewGrocery()
	if err != nil {
		return eris.Wrap(err, "failed to new blacksmith")
	}
	controls.RegisterNPC(bs)

	return nil
}
