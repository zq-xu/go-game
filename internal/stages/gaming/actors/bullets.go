package actors

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/rotisserie/eris"

	"github.com/zq-xu/go-game/internal/stages/gaming/actors/entity"
	"github.com/zq-xu/go-game/pkg/metric"
)

const BulletsName = "Bullets"

type Bullets interface {
	Update(ship *entity.Ship) error
	Draw(screen *ebiten.Image)
	DrawMetrics(screen *ebiten.Image, dc *metric.DrawConfig)

	AddBullet(bullets *entity.Bullet)
	RemoveBullet(bullets *entity.Bullet)

	RangeBullets(fn func(k *entity.Bullet, v bool))
}

type bullets struct {
	lastShootAt time.Time
	bulletss    map[*entity.Bullet]bool

	// TODO move to config
	MaxBulletNum   int           // the max bullets count
	BulletInterval time.Duration // the interval between two bulletss, the unit is ms.
}

func NewBullets() Bullets {
	return &bullets{
		bulletss:       make(map[*entity.Bullet]bool, 0),
		BulletInterval: 200 * time.Millisecond,
		MaxBulletNum:   20,
	}
}

/*
Update: add bullets per interval.
if the bullets is up of the screen, remove it from the storage.
*/
func (s *bullets) Update(ship *entity.Ship) error {
	// if ebiten.IsKeyPressed(ebiten.KeySpace) &&
	// 	len(s.bulletss) < s.MaxBulletNum &&
	// 	time.Since(s.lastShootAt) > s.BulletInterval {
	// 	bullets := entity.NewBullet(s.ctx, ship)
	// 	s.AddBullet(bullets)
	// }

	if time.Since(s.lastShootAt) > s.BulletInterval {
		bullets, err := entity.NewBullet(ship)
		if err != nil {
			return eris.Wrap(err, "new button failed.")
		}
		s.AddBullet(bullets)
	}

	for b := range s.bulletss {
		b.Update()

		if b.IsUpOfScreen() {
			s.RemoveBullet(b)
		}
	}

	return nil
}

func (b *bullets) AddBullet(bullets *entity.Bullet) {
	b.bulletss[bullets] = true
	b.lastShootAt = time.Now()
}

func (b *bullets) RemoveBullet(bullets *entity.Bullet) {
	delete(b.bulletss, bullets)
}

func (b *bullets) Draw(screen *ebiten.Image) {
	for bullets := range b.bulletss {
		bullets.Draw(screen)
	}
}

func (s *bullets) DrawMetrics(screen *ebiten.Image, dc *metric.DrawConfig) {
	op := &text.DrawOptions{}

	op.GeoM.Translate(float64(dc.X), float64(dc.Y))
	op.ColorScale.ScaleWithColor(dc.Color)
	text.Draw(screen, fmt.Sprintf("%s: bulletsCount: %d", BulletsName, len(s.bulletss)), dc.Face, op)

}

func (s *bullets) RangeBullets(fn func(k *entity.Bullet, v bool)) {
	for k, v := range s.bulletss {
		fn(k, v)
	}
}
