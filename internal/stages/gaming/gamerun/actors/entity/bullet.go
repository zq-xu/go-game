package entity

import (
	"github.com/zq-xu/go-game/internal/settings"
	"github.com/zq-xu/go-game/pkg/graphics"
)

const BulletImgPath = "images/bullet.png"

type Bullet struct {
	ImageEntity

	SpeedFactor float64
}

func NewBullet(s *Ship) (*Bullet, error) {
	entity := NewImageEntityWithImage(graphics.GetImage(BulletImgPath),
		settings.GetSettings().ScreenWidth(),
		settings.GetSettings().ScreenHeight())

	entity.UnlimitTop()

	entity.SetX(s.X + float64(s.Img.Width()-entity.Img.Width())/2)
	entity.SetY(s.Y - float64(entity.Img.Height()))

	return &Bullet{
		ImageEntity: *entity,
		SpeedFactor: settings.GetSettings().BulletSpeedFactor(),
	}, nil
}

/*
Update: update the bullet position (based on the speedFactor).
*/
func (b *Bullet) Update() {
	b.MoveUp(b.SpeedFactor)
}
