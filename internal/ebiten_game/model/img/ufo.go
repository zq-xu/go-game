package img

import (
	"math/rand"
	"path"

	"github.com/zq-xu/2d-game/internal/ebiten_game/config"
)

type UFO struct {
	Image

	YSpeedFactor float64
	XSpeedFactor float64
}

func NewUFOImg(cfg *config.Config) *UFO {
	p := path.Join(config.Cfg.ImageRootPath, "ufo.png")

	img := NewImage(p, &cfg.ScreenConfig)

	img.X = float64(rand.Intn(cfg.ScreenWidth-img.Width*2) + img.Width)
	img.Y = -float64(img.Height)

	return &UFO{
		Image:        *img,
		YSpeedFactor: 3,
		XSpeedFactor: 1,
	}
}

/*
Update: update the UFO position (minus speedFactor).
return whether the UFO is out of the screen (for removing it from the storage).
*/
func (u *UFO) Update() bool {
	if u.Y >= float64(u.sc.ScreenHeight) {
		return false
	}

	u.Y += u.YSpeedFactor
	return true
}
