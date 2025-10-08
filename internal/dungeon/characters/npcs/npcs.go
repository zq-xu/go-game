package npcs

import (
	"fmt"
	"image"

	"github.com/rotisserie/eris"
	"github.com/zq-xu/go-game/internal/dungeon/characters/npcs/grocery"
	mageguardian "github.com/zq-xu/go-game/internal/dungeon/characters/npcs/mage_guardian"
	moontower "github.com/zq-xu/go-game/internal/dungeon/characters/npcs/moon_tower"
	"github.com/zq-xu/go-game/internal/dungeon/core/actor"
)

func NewNPCs() ([]actor.Actor, error) {
	list := make([]actor.Actor, 0)
	mg, err := mageguardian.NewMageGuardian(&mageguardian.MageGuardianConfig{
		Name:       fmt.Sprintf("MageGuardian-%s", "blue"),
		StartPoint: &image.Point{X: 700, Y: 300},
	})
	if err != nil {
		return nil, eris.Wrap(err, "failed to new mage guardian")
	}
	list = append(list, mg)

	mt, err := moontower.NewMoonTower(&moontower.MoonTowerConfig{
		Name:       "MoonTower",
		StartPoint: &image.Point{X: 550, Y: 50},
	})
	if err != nil {
		return nil, eris.Wrap(err, "failed to new moon tower")
	}
	list = append(list, mt)

	bs, err := grocery.NewGrocery(&grocery.GroceryConfig{
		Name:       "Grocery",
		StartPoint: &image.Point{X: 100, Y: 650},
	})
	if err != nil {
		return nil, eris.Wrap(err, "failed to new blacksmith")
	}
	list = append(list, bs)
	return list, nil
}
