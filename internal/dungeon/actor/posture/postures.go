package posture

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rotisserie/eris"
	"github.com/zq-xu/go-game/pkg/graphics/images/imagekit"
)

const idleInterval = 3

const (
	UpDirection Direction = iota
	DownDirection
	LeftDirection
	RightDirection
)

const (
	IdleStatus Status = iota
	RunningStatus
	AttackStatus
)

type Status int
type Direction int

type Postures interface {
	UpdateKeyPress(key ebiten.Key)
	UpdateIdle()
	Draw(screen *ebiten.Image, x, y float64)
}

type postures struct {
	status        Status
	movingPosture *movingPosture
	attackPosture *attackPosture
}

func NewPostures() (Postures, error) {
	var err error
	a := &postures{}

	a.movingPosture, err = NewMovingPostures()
	if err != nil {
		return nil, eris.Wrap(err, "new moving postures failed.")
	}

	a.attackPosture, err = NewAttackPosture()
	if err != nil {
		return nil, eris.Wrap(err, "new attack postures failed.")
	}

	return a, nil
}

func (a *postures) UpdateKeyPress(key ebiten.Key) {
	if key == ebiten.KeySpace {
		a.status = AttackStatus
		a.attackPosture.Update(a.movingPosture.currentItem.direction)
		return
	}

	a.movingPosture.UpdateKeyPress(key)
	a.status = a.movingPosture.status
}

func (a *postures) UpdateIdle() {
	a.movingPosture.UpdateIdle()
	a.status = a.movingPosture.status
}

func (a *postures) Draw(screen *ebiten.Image, x, y float64) {
	if a.status == AttackStatus {
		a.attackPosture.item.Draw(screen, x, y)
		// logs.Logger.Debug("direction: ", a.attackPosture.item.direction,
		// 	" current index: ", a.attackPosture.item.rollImages.CurrentIndex())

		return
	}

	a.movingPosture.currentItem.Draw(screen, x, y)
}

func (a *postures) DrawTest(screen *ebiten.Image, x, y float64) {
	a.attackPosture.attackItemSet[DownDirection].rollImages.Range(func(index int, img imagekit.Image) {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(x+float64(50*index), y)
		screen.DrawImage(img.Image(), op)
	})

}
