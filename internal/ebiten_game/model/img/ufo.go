package img

import (
	"math/rand"

	"github.com/zq-xu/2d-game/internal/ebiten_game/brick"
	"github.com/zq-xu/2d-game/internal/ebiten_game/config"
	"github.com/zq-xu/2d-game/internal/ebiten_game/resources"
	"github.com/zq-xu/2d-game/pkg/utils"
)

type UFO struct {
	Image

	YSpeedFactor float64
	XSpeedFactor float64
	XTimesToTop  int

	calX func() float64
}

func NewUFOImg(cfg *config.Config) *UFO {
	img := NewImage(resources.UFOPng, &cfg.ScreenConfig)

	img.X = float64(rand.Intn(cfg.ScreenWidth-img.Width*2) + img.Width)
	img.Y = -float64(img.Height)
	u := &UFO{
		Image:        *img,
		YSpeedFactor: 3,
		XSpeedFactor: 2,
		XTimesToTop:  3 * 60,
	}

	maxTrail := utils.RandomForMinFloat64(img.X, float64(cfg.ScreenWidth)-img.X)
	if maxTrail > 100 {
		u.calX = brick.GenerateRandomStableTrail(img.X, 0, float64(cfg.ScreenHeight), u.XSpeedFactor)
	} else {
		u.calX = brick.GenerateRandomSinTrailWithBase(img.X, maxTrail, u.XTimesToTop)
	}

	return u
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

	u.X = u.calX()
	return true
}
