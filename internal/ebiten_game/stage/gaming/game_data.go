package gaming

import (
	"github.com/zq-xu/2d-game/internal/ebiten_game/entity"
	"github.com/zq-xu/2d-game/internal/ebiten_game/event"
	"github.com/zq-xu/2d-game/internal/ebiten_game/game"
	"github.com/zq-xu/2d-game/pkg/metric"
)

type GameData struct {
	Input *event.Input

	Ship *entity.Ship

	Shoot *Shoot

	UFOs *UFOs

	MetricPool *metric.Pool
}

func NewGameData(ctx *game.Context) *GameData {
	g := &GameData{}

	g.MetricPool = metric.NewMetricPool()
	g.MetricPool.Register("Debug", &metric.DebugerMetrics{})

	g.Input = event.NewInput()
	g.MetricPool.Register(event.InputName, g.Input)

	g.Ship = entity.NewShip(ctx)
	g.MetricPool.Register(entity.ShipName, g.Ship)

	g.Shoot = NewShoot(ctx)
	g.MetricPool.Register(ShootName, g.Shoot)

	g.UFOs = NewUFOs(ctx)
	g.MetricPool.Register(UFOsName, g.UFOs)

	return g
}
