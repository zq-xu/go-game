package actor

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rotisserie/eris"

	"github.com/zq-xu/go-game/internal/dungeon/config"
	"github.com/zq-xu/go-game/internal/dungeon/core/actor/behaviors/attack"
	"github.com/zq-xu/go-game/internal/dungeon/core/actor/behaviors/idle"
	"github.com/zq-xu/go-game/internal/dungeon/core/actor/behaviors/moving"
	"github.com/zq-xu/go-game/pkg/event/collision"
	"github.com/zq-xu/go-game/pkg/event/input"
	"github.com/zq-xu/go-game/pkg/graphics"
	"github.com/zq-xu/go-game/pkg/graphics/images/imagekit"
)

type Actor interface {
	Update()
	Draw(screen *ebiten.Image)

	collision.Object
	MoveDirection(d input.Direction)

	IsAttack() bool
	Attack()
	Idle()
}

type actor struct {
	// collision
	collision.Object

	direction input.Direction

	// status
	status config.ActorStatus

	attack attack.Attack
	idle   idle.Idle
	moving moving.Moving

	op *Option
}

func NewActor(op *Option) (Actor, error) {
	var err error
	a := &actor{
		op:        op,
		direction: input.DefaultDirection,
		status:    config.IdleActorStatus,

		Object: collision.NewResolvRectObject(
			op.Name,
			float64(op.StartPoint.X), float64(op.StartPoint.Y),
			op.Width, op.Height,
			collision.WithStepLength(op.StepLength),
		),
	}

	if len(op.AttackImagePaths) > 0 {
		a.attack, err = attack.NewAttack(op.Name, op.AttackImagePaths)
		if err != nil {
			return nil, eris.Wrap(err, "failed to new actor attack")
		}
	}

	if len(op.IdleImagePaths) > 0 {
		a.idle, err = idle.NewIdle(op.Name, op.IdleImagePaths)
		if err != nil {
			return nil, eris.Wrap(err, "failed to new actor idle")
		}
	}

	if len(op.MovingImagePaths) > 0 {
		a.moving, err = moving.NewMoving(op.Name, a.Object, op.MovingImagePaths)
		if err != nil {
			return nil, eris.Wrap(err, "failed to new actor moving")
		}
	}

	return a, nil
}

func (a *actor) Update() {}

func (a *actor) Draw(screen *ebiten.Image) {
	img := a.image()

	// a.debugCollisionBorder(screen)
	// a.debugImageBorder(screen, img)

	x, y := a.Center()
	leftTopX := x - float64(img.Width())/2
	leftTopY := y - float64(img.Height())/2

	graphics.DrawImage(screen, img.Image(), leftTopX, leftTopY)
}

func (a *actor) image() imagekit.Image {
	switch a.status {
	case config.AttackActorStatus:
		return a.attack.Image()
	case config.RunningActorStatus:
		return a.moving.Image()
	default:
		return a.idle.Image()
	}
}

func (a *actor) IsAttack() bool {
	if a.attack == nil {
		return false
	}

	return a.attack.IsAttack()
}

func (a *actor) Attack() {
	if a.attack == nil {
		return
	}

	a.attack.Attack(a.direction)
	a.status = config.AttackActorStatus
}

func (a *actor) MoveDirection(d input.Direction) {
	if a.moving == nil {
		return
	}

	a.direction = d
	a.moving.Move(a.direction)
	a.status = config.RunningActorStatus
}

func (a *actor) Idle() {
	if a.idle == nil {
		return
	}

	a.idle.Update(a.direction)
	a.status = config.IdleActorStatus
}
