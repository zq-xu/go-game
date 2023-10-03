package gaming

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"

	"github.com/zq-xu/2d-game/internal/ebiten_game/entity"
	"github.com/zq-xu/2d-game/internal/ebiten_game/loader"
	"github.com/zq-xu/2d-game/pkg/metric"
)

const ShootName = "Shoot"

type Shoot struct {
	loader      *loader.Loader
	lastShootAt time.Time
	bullets     map[*entity.Bullet]bool

	// TODO move to config
	MaxBulletNum   int           // the max bullet count
	BulletInterval time.Duration // the interval between two bullets, the unit is ms.
}

func NewShoot(loader *loader.Loader) *Shoot {
	return &Shoot{
		loader:         loader,
		bullets:        make(map[*entity.Bullet]bool, 0),
		BulletInterval: 200 * time.Millisecond,
		MaxBulletNum:   20,
	}
}

func (s *Shoot) Update(ship *entity.Ship) error {
	if ebiten.IsKeyPressed(ebiten.KeySpace) &&
		len(s.bullets) < s.MaxBulletNum &&
		time.Since(s.lastShootAt) > s.BulletInterval {
		bullet := entity.NewBullet(s.loader, ship)
		s.AddBullet(bullet)
	}

	for bullet := range s.bullets {
		inScreen := bullet.Update()
		if !inScreen {
			s.RemoveBullet(bullet)
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

func (s *Shoot) DrawMetrics(screen *ebiten.Image, cfg *metric.DrawConfig) {
	text.Draw(screen, fmt.Sprintf("bulletCount: %d\n", len(s.bullets)), cfg.Face, cfg.X, cfg.Y, cfg.Color)
}

func (s *Shoot) RangeBullets(fn func(k *entity.Bullet, v bool)) {
	for k, v := range s.bullets {
		fn(k, v)
	}
}
