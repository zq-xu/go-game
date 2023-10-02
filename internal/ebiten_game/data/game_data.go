package data

import (
	"github.com/zq-xu/2d-game/internal/ebiten_game/config"
	"github.com/zq-xu/2d-game/internal/ebiten_game/model/img"
	"github.com/zq-xu/2d-game/internal/ebiten_game/model/input"
	"github.com/zq-xu/2d-game/pkg/metric"
)

type GameData struct {
	Cfg *config.Config

	Input *input.Input

	Ship *img.Ship

	Shoot *Shoot

	UFOs *UFOs

	GameStage *GameStage

	MetricPool *metric.Pool
}

func NewGameData() *GameData {
	g := &GameData{
		Cfg:       config.Cfg,
		GameStage: NewGameStage(),
	}

	return g
}

func (g *GameData) Init() {
	g.GameStage = NewGameStage()
	g.MetricPool = metric.NewMetricPool()

	g.Input = input.NewInput()
	g.MetricPool.Register(input.InputName, g.Input)

	g.Ship = img.NewShipImg(config.Cfg)
	g.MetricPool.Register(img.ShipName, g.Ship)

	g.Shoot = NewShoot()
	g.MetricPool.Register(ShootName, g.Shoot)

	g.UFOs = NewUFOs()
	g.MetricPool.Register(UFOsName, g.UFOs)

}

func (g *GameData) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Cfg.ScreenWidth, g.Cfg.ScreenHeight
}
