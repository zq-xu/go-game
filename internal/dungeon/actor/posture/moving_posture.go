package posture

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rotisserie/eris"

	"github.com/zq-xu/go-game/assets/dungeon"
)

var (
	idleImagesPath = map[Direction]string{
		UpDirection:    dungeon.GetDungeonImagePath("/actor/idle/idle_up.png"),
		DownDirection:  dungeon.GetDungeonImagePath("/actor/idle/idle_down.png"),
		LeftDirection:  dungeon.GetDungeonImagePath("/actor/idle/idle_left.png"),
		RightDirection: dungeon.GetDungeonImagePath("/actor/idle/idle_right.png"),
	}

	runImagesPath = map[Direction]string{
		UpDirection:    dungeon.GetDungeonImagePath("/actor/run/run_up.png"),
		DownDirection:  dungeon.GetDungeonImagePath("/actor/run/run_down.png"),
		LeftDirection:  dungeon.GetDungeonImagePath("/actor/run/run_left.png"),
		RightDirection: dungeon.GetDungeonImagePath("/actor/run/run_right.png"),
	}

	keyDirectionSet = map[ebiten.Key]Direction{
		ebiten.KeyDown:  DownDirection,
		ebiten.KeyUp:    UpDirection,
		ebiten.KeyLeft:  LeftDirection,
		ebiten.KeyRight: RightDirection,
	}
)

type movingPosture struct {
	status      Status
	idleCounter int // to change to idles

	currentItem *movingItem

	idleImages map[Direction]*movingItem
	runImages  map[Direction]*movingItem
}

func NewMovingPostures() (*movingPosture, error) {
	var err error
	mp := &movingPosture{
		idleImages: make(map[Direction]*movingItem),
		runImages:  make(map[Direction]*movingItem),
	}

	for k, v := range idleImagesPath {
		mp.idleImages[k], err = newMovingItem(k, v)
		if err != nil {
			return nil, eris.Wrapf(err, "failed to load idle posture item %d", k)
		}
	}

	for k, v := range runImagesPath {
		mp.runImages[k], err = newMovingItem(k, v)
		if err != nil {
			return nil, eris.Wrapf(err, "failed to load run posture item %d", k)
		}
	}

	mp.currentItem = mp.idleImages[DownDirection]
	return mp, nil
}

func (mp *movingPosture) UpdateKeyPress(key ebiten.Key) {
	for k, v := range keyDirectionSet {
		if k == key {
			mp.checkMovingPosture(v)
			return
		}
	}
}

func (mp *movingPosture) checkMovingPosture(d Direction) {
	if mp.currentItem.direction == d && mp.status == RunningStatus {
		mp.currentItem.Update()
		return
	}

	mp.setStep(mp.runImages[d], RunningStatus)
}

func (mp *movingPosture) UpdateIdle() {
	mp.checkIdle()
	mp.currentItem.Update()
}

func (mp *movingPosture) checkIdle() {
	if mp.status == IdleStatus {
		mp.currentItem.Update()
		return
	}

	if mp.idleCounter < idleInterval {
		mp.idleCounter++
		return
	}

	mp.setStep(mp.idleImages[mp.currentItem.direction], IdleStatus)
}

func (mp *movingPosture) setStep(step *movingItem, s Status) {
	mp.currentItem = step

	if s == IdleStatus {
		mp.idleCounter = 0
	}
	mp.status = s
}
