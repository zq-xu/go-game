package entity

import (
	"github.com/zq-xu/2d-game/internal/ebiten_game/loader"
	"github.com/zq-xu/2d-game/pkg/graphics"
)

type Bullet struct {
	graphics.ImageEntity

	SpeedFactor float64
}

func NewBullet(ld *loader.Loader, s *Ship) *Bullet {
	entity := graphics.NewImageEntityWithImage(ld.ImageLoader.GetBulletImage(), ld.Cfg.ScreenWidth, ld.Cfg.ScreenHeight)

	entity.SetX(s.X + float64(s.Img.Width-entity.Img.Width)/2)
	entity.SetY(s.Y - float64(entity.Img.Height))

	return &Bullet{
		ImageEntity: *entity,
		SpeedFactor: 5,
	}
}

/*
Update: update the bullet position (minus speedFactor).
return whether the bullet is out of the screen (for removing it from the storage).
*/
func (b *Bullet) Update() bool {
	if b.IsUpOfScreen() {
		return false
	}

	b.MoveUp(b.SpeedFactor)
	return true
}
