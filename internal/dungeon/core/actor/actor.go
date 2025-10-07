package actor

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rotisserie/eris"

	"github.com/zq-xu/go-game/internal/dungeon/config"
	"github.com/zq-xu/go-game/internal/dungeon/core/actor/behaviors/attack"
	"github.com/zq-xu/go-game/internal/dungeon/core/actor/behaviors/idle"
	"github.com/zq-xu/go-game/internal/dungeon/core/actor/behaviors/running"
	"github.com/zq-xu/go-game/pkg/event/collision"
	"github.com/zq-xu/go-game/pkg/event/input"
	"github.com/zq-xu/go-game/pkg/graphics/images/imagekit"
)

type Actor interface {
	Draw(screen *ebiten.Image)

	collision.Object

	IsAttack() bool
	Attack()
	Moving(key ebiten.Key)
	Idle()
}

type actor struct {
	// collision
	collision.Object

	direction input.Direction

	// status
	status config.ActorStatus

	attack  attack.Attack
	idle    idle.Idle
	running running.Running

	// listen to the key input
	inputListener input.InputListener
}

func NewActor(op *Option) (Actor, error) {
	var err error
	a := &actor{
		direction: input.DefaultDirection,
		status:    config.IdleActorStatus,

		Object: collision.NewResolvRectObject(
			op.Name,
			op.StartX, op.StartY,
			op.Width, op.Height,
		),
	}

	if op.StepLength == 0 {
		op.StepLength = 1
	}

	a.attack, err = attack.NewAttack(op.AttackImagePaths)
	if err != nil {
		return nil, eris.Wrap(err, "failed to new actor attack")
	}

	a.idle, err = idle.NewIdle(op.IdleImagePaths)
	if err != nil {
		return nil, eris.Wrap(err, "failed to new actor idle")
	}

	a.running, err = running.NewRunning(a.Object, op.StepLength, op.MovingImagePaths)
	if err != nil {
		return nil, eris.Wrap(err, "failed to new actor running")
	}

	return a, nil
}

func (a *actor) Draw(screen *ebiten.Image) {
	img := a.image()

	// draw from leftTop
	// keep the leftTop point as default
	leftTopX, leftTopY := a.LeftTop()
	// when the direction is left or top, keep the rightBottom point fixed
	if a.direction == input.LeftDirection ||
		a.direction == input.UpDirection {
		leftTopX = leftTopX + config.ActorWidth - float64(img.Width())
		leftTopY = leftTopY + config.ActorHeight - float64(img.Height())

	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(leftTopX, leftTopY)
	screen.DrawImage(img.Image(), op)
}

func (a *actor) image() imagekit.Image {
	switch a.status {
	case config.AttackActorStatus:
		return a.attack.Image()
	case config.RunningActorStatus:
		return a.running.Image()
	default:
		return a.idle.Image()
	}
}

func (a *actor) IsAttack() bool { return a.attack.IsAttack() }

func (a *actor) Attack() {
	a.attack.Attack(a.direction)
	a.status = config.AttackActorStatus
}

func (a *actor) Moving(key ebiten.Key) {
	a.running.Update(key)
	a.direction = input.GetDirection(key)
	a.status = config.RunningActorStatus
}

func (a *actor) Idle() {
	a.idle.Update(a.direction)
	a.status = config.IdleActorStatus
}
