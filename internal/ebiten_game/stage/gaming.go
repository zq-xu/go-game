package stage

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/2d-game/internal/ebiten_game/entity"
	"github.com/zq-xu/2d-game/internal/ebiten_game/event"
	"github.com/zq-xu/2d-game/internal/ebiten_game/game"
	"github.com/zq-xu/2d-game/internal/ebiten_game/stage/gaming"
)

type GamingStage struct {
	ctx *game.Context

	gameData *gaming.GameData
	status   game.StageStatus
}

func NewGamingStage(ctx *game.Context) *GamingStage {
	return &GamingStage{
		ctx:      ctx,
		gameData: gaming.NewGameData(ctx),
	}
}

func (g *GamingStage) Update() error {
	g.gameData.Input.Update()
	g.gameData.Ship.Update()
	g.gameData.Shoot.Update(g.gameData.Ship)
	g.gameData.UFOs.Update()

	g.checkShootCollision()
	g.checkShipCollision()

	return nil
}

func (g *GamingStage) Draw(screen *ebiten.Image) {
	screen.Fill(g.ctx.Resource.Cfg.BgColor)

	g.gameData.Ship.Draw(screen)
	g.gameData.Input.Draw(screen, g.ctx.Resource.Cfg)
	g.gameData.Shoot.Draw(screen)
	g.gameData.UFOs.Draw(screen)

	g.gameData.MetricPool.DrawMetrics(screen)
}

func (g *GamingStage) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.ctx.Resource.Cfg.Layout(outsideWidth, outsideHeight)
}

func (g *GamingStage) checkShootCollision() {
	g.gameData.Shoot.RangeBullets(func(b *entity.Bullet, bv bool) {
		g.gameData.UFOs.RangeUFOs(func(u *entity.UFO, uv bool) {
			if event.CheckCollision(&u.ImageEntity, &b.ImageEntity) {
				g.gameData.Shoot.RemoveBullet(b)
				g.gameData.UFOs.RemoveUFO(u)
			}
		})
	})
}

func (g *GamingStage) checkShipCollision() {
	g.gameData.UFOs.RangeUFOs(func(u *entity.UFO, uv bool) {
		if event.CheckCollision(&u.ImageEntity, &g.gameData.Ship.ImageEntity) {
			g.status = game.FailStageStatus
		}
	})
}

func (g *GamingStage) GoNextStatus() (bool, Interface) {
	switch g.status {
	case game.SuccessStageStatus, game.FailStageStatus:
		return true, NewEndingStage(g.ctx, g.status)
	default:
		return false, nil
	}
}
