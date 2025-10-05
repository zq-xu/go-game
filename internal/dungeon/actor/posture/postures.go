package posture

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rotisserie/eris"

	"github.com/zq-xu/go-game/internal/dungeon/actor/posture/attack"
	"github.com/zq-xu/go-game/internal/dungeon/actor/posture/moving"
	"github.com/zq-xu/go-game/internal/dungeon/config"
)

type Postures interface {
	UpdateKeyPress(key ebiten.Key)

	Draw(screen *ebiten.Image, x, y float64)

	IsAttacking() bool
}

type postures struct {
	status        config.ActorStatus
	movingPosture moving.MovingkPosture
	attackPosture attack.AttackPosture
}

func NewPostures() (Postures, error) {
	var err error
	a := &postures{}

	a.movingPosture, err = moving.NewMovingPostures()
	if err != nil {
		return nil, eris.Wrap(err, "new moving postures failed.")
	}

	a.attackPosture, err = attack.NewGifAttackPosture()
	if err != nil {
		return nil, eris.Wrap(err, "new attack postures failed.")
	}

	return a, nil
}

func (a *postures) UpdateKeyPress(key ebiten.Key) {
	if a.attackPosture.IsAttacking() {
		return
	}

	if key == ebiten.KeySpace {
		a.status = config.AttackActorStatus
		a.attackPosture.Attack(a.movingPosture.Direction())
		return
	}

	a.movingPosture.UpdateKeyPress(key)
	a.status = a.movingPosture.Status()
}

func (a *postures) Draw(screen *ebiten.Image, x, y float64) {
	if a.status == config.AttackActorStatus {
		a.attackPosture.Draw(screen, x, y)
		// logs.Logger.Debug("direction: ", a.attackPosture.item.direction,
		// 	" current index: ", a.attackPosture.item.rollImages.CurrentIndex())

		return
	}

	a.movingPosture.Draw(screen, x, y)
}

func (a *postures) IsAttacking() bool {
	return a.attackPosture.IsAttacking()
}
