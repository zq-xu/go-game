package gameui

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/internal/ebiten_game/entity"
	"github.com/zq-xu/go-game/internal/ebiten_game/event"
	"github.com/zq-xu/go-game/internal/ebiten_game/game"
	"github.com/zq-xu/go-game/internal/ebiten_game/stage/gameui/runtime"
	"github.com/zq-xu/go-game/pkg/metric"
)

const runtimeName = "Runtime"

type RuntimeOpt func(r *Runtime)

type Runtime struct {
	ctx *game.Context

	Ship *entity.Ship

	Shoot *runtime.Shoot

	UFOs *runtime.UFOs

	MetricPool *metric.Pool

	Status game.Status

	collissionCallback func(bool)
}

func WithRuntimeCollission(callback func(bool)) RuntimeOpt {
	return func(r *Runtime) {
		r.collissionCallback = callback
	}
}

func NewRuntime(ctx *game.Context, opts ...RuntimeOpt) *Runtime {
	g := &Runtime{ctx: ctx}

	for _, fn := range opts {
		fn(g)
	}

	g.MetricPool = metric.NewMetricPool()
	metric.MultiPool.Add(runtimeName, g.MetricPool)

	g.Ship = entity.NewShip(ctx)
	g.MetricPool.Register(entity.ShipName, g.Ship)

	g.Shoot = runtime.NewShoot(ctx)
	g.MetricPool.Register(runtime.ShootName, g.Shoot)

	g.UFOs = runtime.NewUFOs(ctx)
	g.MetricPool.Register(runtime.UFOsName, g.UFOs)

	return g
}

func (g *Runtime) Update() error {
	g.Ship.Update()
	g.Shoot.Update(g.Ship)
	g.UFOs.Update()

	g.checkShootCollision()
	g.checkShipCollision()

	return nil
}

func (g *Runtime) Draw(screen *ebiten.Image) {
	g.Ship.Draw(screen)
	g.Shoot.Draw(screen)
	g.UFOs.Draw(screen)
}

func (g *Runtime) checkShootCollision() {
	g.Shoot.RangeBullets(func(b *entity.Bullet, bv bool) {
		g.UFOs.RangeUFOs(func(u *entity.UFO, uv bool) {
			if event.CheckCollision(&u.ImageEntity, &b.ImageEntity) {
				g.Shoot.RemoveBullet(b)
				g.UFOs.RemoveUFO(u)
			}
		})
	})
}

func (g *Runtime) checkShipCollision() {
	if g.collissionCallback == nil {
		return
	}

	g.UFOs.RangeUFOs(func(u *entity.UFO, uv bool) {
		g.collissionCallback(event.CheckCollision(&u.ImageEntity, &g.Ship.ImageEntity))
	})
}
