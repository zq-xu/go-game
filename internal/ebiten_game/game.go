package ebitengame

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/2d-game/internal/ebiten_game/config"
	"github.com/zq-xu/2d-game/internal/ebiten_game/metric"
	"github.com/zq-xu/2d-game/internal/ebiten_game/model/img"
	"github.com/zq-xu/2d-game/internal/ebiten_game/model/input"
)

type Game struct {
	cfg *config.Config

	input *input.Input

	ship *img.Ship

	shoot *Shoot
}

func NewGame() *Game {
	ebiten.SetWindowSize(config.Cfg.ScreenWidth, config.Cfg.ScreenHeight)
	ebiten.SetWindowTitle(config.Cfg.Title)

	g := &Game{
		input: input.NewInput(),
		cfg:   config.Cfg,
		ship:  img.NewShipImg(config.Cfg),
		shoot: NewShoot(),
	}

	metric.Register(input.InputName, g.input)
	metric.Register(img.ShipName, g.ship)
	metric.Register(ShootName, g.shoot)

	return g
}

func (g *Game) Update() error {
	g.input.Update()
	g.ship.Update()
	g.shoot.Update(g.cfg, g.ship)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(g.cfg.BgColor)

	g.ship.Draw(screen)
	g.input.Draw(screen, g.cfg)
	g.shoot.Draw(screen)

	metric.DrawMetrics(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.cfg.ScreenWidth, g.cfg.ScreenHeight
}
