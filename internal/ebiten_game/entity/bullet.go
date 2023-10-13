package entity

import (
	"github.com/zq-xu/go-game/internal/ebiten_game/game"
	"github.com/zq-xu/go-game/internal/ebiten_game/resource/loader"
	"github.com/zq-xu/go-game/pkg/graphics"
)

type Bullet struct {
	graphics.ImageEntity

	SpeedFactor float64
}

func NewBullet(ctx *game.Context, s *Ship) *Bullet {
	img := ctx.Resource.ImageLoader.ImgLoader.MustGetImage(loader.BulletImgPath)
	entity := graphics.NewImageEntityWithImage(img, ctx.Resource.ScreenWidth, ctx.Resource.ScreenHeight)

	entity.UnlimitTop()

	entity.SetX(s.X + float64(s.Img.Width-entity.Img.Width)/2)
	entity.SetY(s.Y - float64(entity.Img.Height))

	return &Bullet{
		ImageEntity: *entity,
		SpeedFactor: ctx.Resource.BulletSpeedFactor,
	}
}

/*
Update: update the bullet position (based on the speedFactor).
*/
func (b *Bullet) Update() {
	b.MoveUp(b.SpeedFactor)
}
