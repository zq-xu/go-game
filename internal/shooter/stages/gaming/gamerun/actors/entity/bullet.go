package entity

import (
	"github.com/rotisserie/eris"
	"github.com/zq-xu/go-game/assets"
	"github.com/zq-xu/go-game/internal/shooter/settings"
)

var BulletImgPath = assets.GetShooterImagePath("bullet.png")

type Bullet struct {
	ImageEntity

	SpeedFactor float64
}

func NewBullet(s *Ship) (*Bullet, error) {
	entity, err := NewImageEntity(BulletImgPath,
		settings.GetSettings().ScreenWidth(),
		settings.GetSettings().ScreenHeight())
	if err != nil {
		return nil, eris.Wrap(err, "new bullet image entity failed")
	}
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
