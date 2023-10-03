package img

import (
	"github.com/zq-xu/2d-game/internal/ebiten_game/config"
	"github.com/zq-xu/2d-game/internal/ebiten_game/resources"
)

type Bullet struct {
	Image

	SpeedFactor float64
}

func NewBulletImg(cfg *config.Config, s *Ship) *Bullet {
	img := NewImage(resources.BulletPng, &cfg.ScreenConfig)

	img.X = s.X + float64(s.Width-img.Width)/2
	img.Y = s.Y - float64(img.Height)

	return &Bullet{
		Image:       *img,
		SpeedFactor: 5,
	}
}

/*
Update: update the bullet position (minus speedFactor).
return whether the bullet is out of the screen (for removing it from the storage).
*/
func (b *Bullet) Update() bool {
	if b.Y <= -float64(b.Width) {
		return false
	}

	b.Y -= b.SpeedFactor
	return true
}
