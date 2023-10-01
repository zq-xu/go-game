package mode

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/2d-game/internal/ebiten_game/data"
	"github.com/zq-xu/2d-game/internal/ebiten_game/event"
	"github.com/zq-xu/2d-game/internal/ebiten_game/model/img"
)

type GamingModeHandler struct {
	gameData *data.GameData
	ls       ModeListener
}

func init() {
	Register(GamingMode, NewGamingModeHandler)
}

func NewGamingModeHandler(gameData *data.GameData, ls ModeListener) GameModeHandler {
	gameData.Init()

	return &GamingModeHandler{
		gameData: gameData,
		ls:       ls,
	}
}

func (g *GamingModeHandler) Update() error {
	g.gameData.Input.Update()
	g.gameData.Ship.Update()
	g.gameData.Shoot.Update(g.gameData.Cfg, g.gameData.Ship)
	g.gameData.UFOs.Update(g.gameData.Cfg)

	g.checkShootCollision()
	g.checkShipCollision()

	return nil
}

func (g *GamingModeHandler) Draw(screen *ebiten.Image) {
	screen.Fill(g.gameData.Cfg.BgColor)

	g.gameData.Ship.Draw(screen)
	g.gameData.Input.Draw(screen, g.gameData.Cfg)
	g.gameData.Shoot.Draw(screen)
	g.gameData.UFOs.Draw(screen)

	g.gameData.MetricPool.DrawMetrics(screen)
}

func (g *GamingModeHandler) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.gameData.Layout(outsideWidth, outsideHeight)
}

func (g *GamingModeHandler) checkShootCollision() {
	g.gameData.Shoot.RangeBullets(func(b *img.Bullet, bv bool) {
		g.gameData.UFOs.RangeUFOs(func(u *img.UFO, uv bool) {
			if event.CheckCollision(&u.Image, &b.Image) {
				g.gameData.Shoot.RemoveBullet(b)
				g.gameData.UFOs.RemoveUFO(u)
			}
		})
	})
}

func (g *GamingModeHandler) checkShipCollision() {
	g.gameData.UFOs.RangeUFOs(func(u *img.UFO, uv bool) {
		if event.CheckCollision(&u.Image, &g.gameData.Ship.Image) {
			g.gameData.GameStage.Result = data.FailStageResult
			g.ls.SetMode(GameOverMode)
		}
	})
}
