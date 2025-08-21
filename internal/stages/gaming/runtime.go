package gaming

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/internal/listener"
	"github.com/zq-xu/go-game/internal/stages/gaming/actors"
	"github.com/zq-xu/go-game/internal/stages/gaming/actors/entity"
	"github.com/zq-xu/go-game/internal/status"
	"github.com/zq-xu/go-game/pkg/metric"
)

const actorsName = "Runtime"

type RuntimeOpt func(r *Runtime)

type Runtime struct {
	Ship    *actors.Ship
	Bullets actors.Bullets
	UFOs    *actors.UFOs

	MetricPool         *metric.Pool
	Status             status.Status
	collissionCallback func(bool)
}

func WithRuntimeCollission(callback func(bool)) RuntimeOpt {
	return func(r *Runtime) {
		r.collissionCallback = callback
	}
}

func NewRuntime(opts ...RuntimeOpt) *Runtime {
	g := &Runtime{}

	for _, fn := range opts {
		fn(g)
	}

	g.MetricPool = metric.NewMetricPool()
	metric.MultiPool.Add(actorsName, g.MetricPool)

	g.Ship = actors.NewShip()
	g.MetricPool.Register(actors.ShipName, g.Ship)

	g.Bullets = actors.NewBullets()
	g.MetricPool.Register(actors.BulletsName, g.Bullets)

	g.UFOs = actors.NewUFOs()
	g.MetricPool.Register(actors.UFOsName, g.UFOs)

	return g
}

func (g *Runtime) Update() error {
	g.Ship.Update()
	g.Bullets.Update(g.Ship.Ship)
	g.UFOs.Update()

	g.checkBulletsCollision()
	g.checkShipCollision()

	return nil
}

func (g *Runtime) Draw(screen *ebiten.Image) {
	g.Ship.Draw(screen)
	g.Bullets.Draw(screen)
	g.UFOs.Draw(screen)
}

func (g *Runtime) checkBulletsCollision() {
	g.Bullets.RangeBullets(func(b *entity.Bullet, bv bool) {
		g.UFOs.RangeUFOs(func(u *entity.UFO, uv bool) {
			if actors.CheckCollision(&u.ImageEntity, &b.ImageEntity) {
				g.Bullets.RemoveBullet(b)
				g.UFOs.RemoveUFO(u)
				listener.GetListener().GameDataListener().AddShotUFO(1)
			}
		})
	})
}

func (g *Runtime) checkShipCollision() {
	if g.collissionCallback == nil {
		return
	}

	g.UFOs.RangeUFOs(func(u *entity.UFO, uv bool) {
		g.collissionCallback(actors.CheckCollision(&u.ImageEntity, &g.Ship.ImageEntity))
	})
}
