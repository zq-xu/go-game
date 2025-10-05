package moving

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rotisserie/eris"

	"github.com/zq-xu/go-game/assets/dungeon"
	"github.com/zq-xu/go-game/internal/dungeon/config"
)

const idleInterval = 3

var (
	idleImagesPath = map[config.Direction]string{
		config.UpDirection:    dungeon.GetDungeonImagePath("/actor/idle/idle_up.png"),
		config.DownDirection:  dungeon.GetDungeonImagePath("/actor/idle/idle_down.png"),
		config.LeftDirection:  dungeon.GetDungeonImagePath("/actor/idle/idle_left.png"),
		config.RightDirection: dungeon.GetDungeonImagePath("/actor/idle/idle_right.png"),
	}

	runImagesPath = map[config.Direction]string{
		config.UpDirection:    dungeon.GetDungeonImagePath("/actor/run/run_up.png"),
		config.DownDirection:  dungeon.GetDungeonImagePath("/actor/run/run_down.png"),
		config.LeftDirection:  dungeon.GetDungeonImagePath("/actor/run/run_left.png"),
		config.RightDirection: dungeon.GetDungeonImagePath("/actor/run/run_right.png"),
	}

	keyDirectionSet = map[ebiten.Key]config.Direction{
		ebiten.KeyDown:  config.DownDirection,
		ebiten.KeyUp:    config.UpDirection,
		ebiten.KeyLeft:  config.LeftDirection,
		ebiten.KeyRight: config.RightDirection,
	}
)

type MovingkPosture interface {
	UpdateKeyPress(key ebiten.Key)

	Draw(screen *ebiten.Image, x, y float64)

	Direction() config.Direction
	Status() config.ActorStatus
}

type movingPosture struct {
	status      config.ActorStatus
	idleCounter int // to change to idles

	currentItem *movingItem

	idleImages map[config.Direction]*movingItem
	runImages  map[config.Direction]*movingItem
}

func NewMovingPostures() (MovingkPosture, error) {
	var err error
	mp := &movingPosture{
		idleImages: make(map[config.Direction]*movingItem),
		runImages:  make(map[config.Direction]*movingItem),
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

	mp.currentItem = mp.idleImages[config.DefaultDirection]
	return mp, nil
}

func (mp *movingPosture) UpdateKeyPress(key ebiten.Key) {
	mp.checkStatus(key)
	mp.currentItem.Update()
}

func (mp *movingPosture) checkStatus(key ebiten.Key) {
	for k, v := range keyDirectionSet {
		if k == key {
			mp.checkMovingPosture(v)
			return
		}
	}

	mp.checkIdle()
}

func (mp *movingPosture) checkMovingPosture(d config.Direction) {
	if mp.currentItem.direction == d && mp.status == config.RunningActorStatus {
		return
	}

	mp.setStep(mp.runImages[d], config.RunningActorStatus)
}

func (mp *movingPosture) checkIdle() {
	if mp.status == config.IdleActorStatus {
		mp.currentItem.Update()
		return
	}

	if mp.idleCounter < idleInterval {
		mp.idleCounter++
		return
	}

	mp.setStep(mp.idleImages[mp.currentItem.direction], config.IdleActorStatus)
}

func (mp *movingPosture) setStep(step *movingItem, s config.ActorStatus) {
	mp.currentItem = step

	if s == config.IdleActorStatus {
		mp.idleCounter = 0
	}
	mp.status = s
}

func (mp *movingPosture) Draw(screen *ebiten.Image, x, y float64) {
	mp.currentItem.Draw(screen, x, y)
}

func (mp *movingPosture) Direction() config.Direction { return mp.currentItem.direction }
func (mp *movingPosture) Status() config.ActorStatus  { return mp.status }
