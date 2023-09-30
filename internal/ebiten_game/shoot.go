package ebitengame

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"

	"github.com/zq-xu/2d-game/internal/ebiten_game/config"
	"github.com/zq-xu/2d-game/internal/ebiten_game/metric"
	"github.com/zq-xu/2d-game/internal/ebiten_game/model/img"
)

const ShootName = "Shoot"

type Shoot struct {
	lastShootAt time.Time
	bullets     map[*img.Bullet]bool
}

func NewShoot() *Shoot {
	return &Shoot{bullets: make(map[*img.Bullet]bool, 0)}
}

func (s *Shoot) Update(cfg *config.Config, ship *img.Ship) error {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		bullet := img.NewBulletImg(cfg, ship)
		s.addBullet(bullet)
	}

	for bullet := range s.bullets {
		inScreen := bullet.Update()
		if !inScreen {
			s.removeBullet(bullet)
		}
	}

	return nil
}

func (s *Shoot) addBullet(bullet *img.Bullet) {
	if time.Since(s.lastShootAt) < time.Millisecond*200 {
		return
	}

	s.bullets[bullet] = true
	s.lastShootAt = time.Now()
}

func (s *Shoot) removeBullet(bullet *img.Bullet) {
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
