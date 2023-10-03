package gaming

import (
	"github.com/zq-xu/2d-game/internal/ebiten_game/entity"
	"github.com/zq-xu/2d-game/internal/ebiten_game/event"
	"github.com/zq-xu/2d-game/internal/ebiten_game/loader"
	"github.com/zq-xu/2d-game/pkg/metric"
)

type GameData struct {
	Input *event.Input

	Ship *entity.Ship

	Shoot *Shoot

	UFOs *UFOs

	MetricPool *metric.Pool
}

func NewGameData(loader *loader.Loader) *GameData {
	g := &GameData{}

	g.MetricPool = metric.NewMetricPool()

	g.Input = event.NewInput()
	g.MetricPool.Register(event.InputName, g.Input)

	g.Ship = entity.NewShip(loader)
	g.MetricPool.Register(entity.ShipName, g.Ship)

	g.Shoot = NewShoot(loader)
	g.MetricPool.Register(ShootName, g.Shoot)

	g.UFOs = NewUFOs(loader)
	g.MetricPool.Register(UFOsName, g.UFOs)

	return g
}
