package runtime

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"

	"github.com/zq-xu/2d-game/internal/ebiten_game/entity"
	"github.com/zq-xu/2d-game/internal/ebiten_game/game"
	"github.com/zq-xu/2d-game/pkg/metric"
)

const ShootName = "Shoot"

type Shoot struct {
	ctx *game.Context

	lastShootAt time.Time
	bullets     map[*entity.Bullet]bool

	// TODO move to config
	MaxBulletNum   int           // the max bullet count
	BulletInterval time.Duration // the interval between two bullets, the unit is ms.
}

func NewShoot(ctx *game.Context) *Shoot {
	return &Shoot{
		ctx:            ctx,
		bullets:        make(map[*entity.Bullet]bool, 0),
		BulletInterval: 200 * time.Millisecond,
		MaxBulletNum:   20,
	}
}

/*
Update: add bullet per interval.
if the bullet is up of the screen, remove it from the storage.
*/
func (s *Shoot) Update(ship *entity.Ship) error {
	// if ebiten.IsKeyPressed(ebiten.KeySpace) &&
	// 	len(s.bullets) < s.MaxBulletNum &&
	// 	time.Since(s.lastShootAt) > s.BulletInterval {
	// 	bullet := entity.NewBullet(s.ctx, ship)
	// 	s.AddBullet(bullet)
	// }

	if time.Since(s.lastShootAt) > s.BulletInterval {
		bullet := entity.NewBullet(s.ctx, ship)
		s.AddBullet(bullet)
	}

	for b := range s.bullets {
		b.Update()

		if b.IsUpOfScreen() {
			s.RemoveBullet(b)
		}
	}

	return nil
}

func (s *Shoot) AddBullet(bullet *entity.Bullet) {
	s.bullets[bullet] = true
	s.lastShootAt = time.Now()
}

func (s *Shoot) RemoveBullet(bullet *entity.Bullet) {
	delete(s.bullets, bullet)
}

func (s *Shoot) Draw(screen *ebiten.Image) {
	for bullet := range s.bullets {
		bullet.Draw(screen)
	}
}

func (s *Shoot) DrawMetrics(screen *ebiten.Image, dc *metric.DrawConfig) {
	text.Draw(screen, fmt.Sprintf("%s: ", ShootName), dc.Face, dc.X, dc.Y, dc.Color)
	text.Draw(screen,
		fmt.Sprintf("bulletCount: %d\n", len(s.bullets)),
		dc.Face,
		dc.X+metric.MetricCharWidth*(len(ShootName)+5),
		dc.Y,
		dc.Color,
	)
}

func (s *Shoot) RangeBullets(fn func(k *entity.Bullet, v bool)) {
	for k, v := range s.bullets {
		fn(k, v)
	}
}
