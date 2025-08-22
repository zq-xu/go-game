package gamerun

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/internal/listener"
	"github.com/zq-xu/go-game/internal/stages/gaming/gamerun/actors"
	"github.com/zq-xu/go-game/internal/stages/gaming/gamerun/actors/entity"
	"github.com/zq-xu/go-game/internal/status"
	"github.com/zq-xu/go-game/pkg/metric"
)

const actorsName = "GameRun"

type GameRun interface {
	Update() error
	Draw(screen *ebiten.Image)
}

type gameRunOpt func(r *gameRun)

type gameRun struct {
	Ship    *actors.Ship
	Bullets actors.Bullets
	UFOs    *actors.UFOs

	MetricPool         *metric.Pool
	Status             status.Status
	collissionCallback func(bool)
}

func WithgameRunCollission(callback func(bool)) gameRunOpt {
	return func(r *gameRun) {
		r.collissionCallback = callback
	}
}

func NewGameRun(opts ...gameRunOpt) GameRun {
	g := &gameRun{}

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

func (g *gameRun) Update() error {
	g.Ship.Update()
	g.Bullets.Update(g.Ship.Ship)
	g.UFOs.Update()

	g.checkBulletsCollision()
	g.checkShipCollision()

	return nil
}

func (g *gameRun) Draw(screen *ebiten.Image) {
	g.Ship.Draw(screen)
	g.Bullets.Draw(screen)
	g.UFOs.Draw(screen)
}

func (g *gameRun) checkBulletsCollision() {
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

func (g *gameRun) checkShipCollision() {
	if g.collissionCallback == nil {
		return
	}

	g.UFOs.RangeUFOs(func(u *entity.UFO, uv bool) {
		g.collissionCallback(actors.CheckCollision(&u.ImageEntity, &g.Ship.ImageEntity))
	})
}
