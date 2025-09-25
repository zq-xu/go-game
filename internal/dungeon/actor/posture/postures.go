package posture

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rotisserie/eris"

	"github.com/zq-xu/go-game/assets"
)

const idleInterval = 3

const (
	UpDirection Direction = iota
	DownDirection
	LeftDirection
	RightDirection
)

var (
	idleImagesPath = map[Direction]string{
		UpDirection:    assets.GetDungeonImagePath("/actor/idle/idle_up.png"),
		DownDirection:  assets.GetDungeonImagePath("/actor/idle/idle_down.png"),
		LeftDirection:  assets.GetDungeonImagePath("/actor/idle/idle_left.png"),
		RightDirection: assets.GetDungeonImagePath("/actor/idle/idle_right.png"),
	}

	runImagesPath = map[Direction]string{
		UpDirection:    assets.GetDungeonImagePath("/actor/run/run_up.png"),
		DownDirection:  assets.GetDungeonImagePath("/actor/run/run_down.png"),
		LeftDirection:  assets.GetDungeonImagePath("/actor/run/run_left.png"),
		RightDirection: assets.GetDungeonImagePath("/actor/run/run_right.png"),
	}

	keyDirectionSet = map[ebiten.Key]Direction{
		ebiten.KeyDown:  DownDirection,
		ebiten.KeyUp:    UpDirection,
		ebiten.KeyLeft:  LeftDirection,
		ebiten.KeyRight: RightDirection,
	}
)

type Direction int

type Postures interface {
	Update()
	Draw(screen *ebiten.Image, x, y float64)
}

type postures struct {
	idleCounter int // to change to idles
	isIdle      bool

	direction   Direction
	currentStep postureItem

	idleImages map[Direction]postureItem
	runImages  map[Direction]postureItem
}

func NewPostures() (Postures, error) {
	var err error
	a := &postures{
		idleImages: make(map[Direction]postureItem),
		runImages:  make(map[Direction]postureItem),
	}

	for k, v := range idleImagesPath {
		a.idleImages[k], err = newPostureItem(v)
		if err != nil {
			return nil, eris.Wrapf(err, "failed to load posture item %d", k)
		}
	}

	for k, v := range runImagesPath {
		a.runImages[k], err = newPostureItem(v)
		if err != nil {
			return nil, eris.Wrapf(err, "failed to load posture ite %d", k)
		}
	}

	a.currentStep = a.idleImages[DownDirection]
	return a, nil
}

func (a *postures) Update() {
	a.updatePosture()
	a.currentStep.Update()
}

func (a *postures) updatePosture() {
	for k, v := range keyDirectionSet {
		if ebiten.IsKeyPressed(k) {
			a.moving(v)
			return
		}
	}

	a.checkIdle()
}

func (a *postures) moving(d Direction) {
	if a.direction == d && !a.isIdle {
		return
	}

	a.isIdle = false
	a.idleCounter = 0

	a.direction = d
	a.currentStep = a.runImages[d]
}

func (a *postures) checkIdle() {
	if a.isIdle {
		return
	}

	if a.idleCounter < idleInterval {
		a.idleCounter++
		return
	}

	a.isIdle = true
	a.idleCounter = 0
	a.currentStep = a.idleImages[a.direction]
}

func (a *postures) Draw(screen *ebiten.Image, x, y float64) {
	a.currentStep.Draw(screen, x, y)
}
