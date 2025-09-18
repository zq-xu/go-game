package gamerun

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/internal/shooter/data"
	"github.com/zq-xu/go-game/internal/shooter/stages"
	"github.com/zq-xu/go-game/internal/shooter/stages/gaming/gamerun/actors"
	"github.com/zq-xu/go-game/internal/shooter/stages/gaming/gamerun/actors/entity"
	"github.com/zq-xu/go-game/internal/shooter/status"
	"github.com/zq-xu/go-game/internal/shooter/ui/components"
	"github.com/zq-xu/go-game/pkg/metrics"
)

const actorsName = "GameRun"

type GameRun interface {
	Update()
	Draw(screen *ebiten.Image)
}

type gameRun struct {
	ctx      stages.StageContext
	gameData data.Data

	background ebiten.Game

	ship    *actors.Ship
	bullets actors.Bullets
	ufos    *actors.UFOs
}

func NewGameRun(ctx stages.StageContext, gameData data.Data) GameRun {
	g := &gameRun{
		ctx:        ctx,
		gameData:   gameData,
		background: components.NewDeepStarrySkyDownwardsBackground(),
	}

	metricPool := metrics.NewMetricPool()
	gameData.Metrics().Add(actorsName, metricPool)

	g.ship = actors.NewShip()
	metricPool.Register(actors.ShipName, g.ship)

	g.bullets = actors.NewBullets()
	metricPool.Register(actors.BulletsName, g.bullets)

	g.ufos = actors.NewUFOs()
	metricPool.Register(actors.UFOsName, g.ufos)

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
				g.gameData.AddShotUFO(1)
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
