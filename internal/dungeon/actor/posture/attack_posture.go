package posture

import (
	"github.com/rotisserie/eris"

	"github.com/zq-xu/go-game/assets/dungeon"
)

var (
	attackImagesPath = map[Direction]string{
		UpDirection:    dungeon.GetDungeonImagePath("/actor/attack/attack1_up.png"),
		DownDirection:  dungeon.GetDungeonImagePath("/actor/attack/attack1_down.png"),
		LeftDirection:  dungeon.GetDungeonImagePath("/actor/attack/attack1_left.png"),
		RightDirection: dungeon.GetDungeonImagePath("/actor/attack/attack1_right.png"),
	}
)

type attackPosture struct {
	item          *attackItem
	attackItemSet map[Direction]*attackItem
}

func NewAttackPosture() (*attackPosture, error) {
	var err error
	ap := &attackPosture{
		attackItemSet: make(map[Direction]*attackItem),
	}

	for k, v := range attackImagesPath {
		ap.attackItemSet[k], err = newAttackItem(k, v)
		if err != nil {
			return nil, eris.Wrapf(err, "failed to load attack posture item %d", k)
		}
	}

	ap.item = ap.attackItemSet[DownDirection]
	return ap, nil
}

func (ap *attackPosture) refreshItemByDirection(d Direction) {
	if d == ap.item.direction {
		ap.item.Update()
		return
	}

	i := ap.attackItemSet[d]
	ap.item = i
	ap.item.MoveToStart()
}

func (ap *attackPosture) Update(direction Direction) {
	ap.refreshItemByDirection(direction)
	// logs.Logger.Debug("direction: ", ap.item.direction, " current index: ", ap.item.rollImages.CurrentIndex())
}
