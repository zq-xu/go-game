package gamerun

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/internal/listener"
	"github.com/zq-xu/go-game/internal/stages"
	"github.com/zq-xu/go-game/internal/stages/gaming/gamerun/actors"
	"github.com/zq-xu/go-game/internal/stages/gaming/gamerun/actors/entity"
	"github.com/zq-xu/go-game/internal/status"
	"github.com/zq-xu/go-game/internal/ui/components"
	"github.com/zq-xu/go-game/pkg/metric"
)

const actorsName = "GameRun"

type GameRun interface {
	Update()
	Draw(screen *ebiten.Image)
}

type gameRun struct {
	ctx stages.StageContext

	ship    *actors.Ship
	bullets actors.Bullets
	ufos    *actors.UFOs

	background ebiten.Game

	MetricPool *metric.Pool
	Status     status.Status
}

func NewGameRun(ctx stages.StageContext) GameRun {
	g := &gameRun{}

	g.MetricPool = metric.NewMetricPool()
	metric.MultiPool.Add(actorsName, g.MetricPool)

	g.ship = actors.NewShip()
	g.MetricPool.Register(actors.ShipName, g.ship)

	g.bullets = actors.NewBullets()
	g.MetricPool.Register(actors.BulletsName, g.bullets)

	g.ufos = actors.NewUFOs()
	g.MetricPool.Register(actors.UFOsName, g.ufos)

	g.background = components.NewDeepStarrySkyDownwardsBackground()
	return g
}

func (g *gameRun) Update() {
	g.ship.Update()
	g.bullets.Update(g.ship.Ship)
	g.ufos.Update()

	g.checkBulletsCollision()
	g.checkShipCollision()

	g.background.Update()
}

func (g *gameRun) Draw(screen *ebiten.Image) {
	g.background.Draw(screen)

	g.ship.Draw(screen)
	g.bullets.Draw(screen)
	g.ufos.Draw(screen)
}

func (g *gameRun) checkBulletsCollision() {
	g.bullets.RangeBullets(func(b *entity.Bullet, bv bool) {
		g.ufos.RangeUFOs(func(u *entity.UFO, uv bool) {
			if actors.CheckCollision(&u.ImageEntity, &b.ImageEntity) {
				g.bullets.RemoveBullet(b)
				g.ufos.RemoveUFO(u)
				listener.GetListener().GameDataListener().AddShotUFO(1)
			}
		})
	})
}

func (g *gameRun) checkShipCollision() {
	g.ufos.RangeUFOs(func(u *entity.UFO, uv bool) {
		b := actors.CheckCollision(&u.ImageEntity, &g.ship.ImageEntity)
		if !b {
			return
		}
		g.ctx.SetStatus(status.FailStatus)
		g.ctx.SetCurrentGameStage(stages.EndingStage)
	})
}
