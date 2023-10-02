package data

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"

	"github.com/zq-xu/2d-game/internal/ebiten_game/config"
	"github.com/zq-xu/2d-game/internal/ebiten_game/model/img"
	"github.com/zq-xu/2d-game/pkg/metric"
)

const ShootName = "Shoot"

type Shoot struct {
	lastShootAt time.Time
	bullets     map[*img.Bullet]bool

	// TODO move to config
	MaxBulletNum   int           // the max bullet count
	BulletInterval time.Duration // the interval between two bullets, the unit is ms.
}

func NewShoot() *Shoot {
	return &Shoot{
		bullets:        make(map[*img.Bullet]bool, 0),
		BulletInterval: 200 * time.Millisecond,
		MaxBulletNum:   20,
	}
}

func (s *Shoot) Update(cfg *config.Config, ship *img.Ship) error {
	if ebiten.IsKeyPressed(ebiten.KeySpace) &&
		len(s.bullets) < s.MaxBulletNum &&
		time.Since(s.lastShootAt) > s.BulletInterval {
		bullet := img.NewBulletImg(cfg, ship)
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

func (s *Shoot) AddBullet(bullet *img.Bullet) {
	s.bullets[bullet] = true
	s.lastShootAt = time.Now()
}

func (s *Shoot) RemoveBullet(bullet *img.Bullet) {
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

func (s *Shoot) RangeBullets(fn func(k *img.Bullet, v bool)) {
	for k, v := range s.bullets {
		fn(k, v)
	}
}
