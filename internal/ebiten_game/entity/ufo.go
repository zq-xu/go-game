package entity

import (
	"math/rand"

	"github.com/zq-xu/2d-game/internal/ebiten_game/loader"
	"github.com/zq-xu/2d-game/pkg/brick"
	"github.com/zq-xu/2d-game/pkg/graphics"
	"github.com/zq-xu/2d-game/pkg/utils"
)

type UFO struct {
	graphics.ImageEntity

	YSpeedFactor float64
	XSpeedFactor float64
	XTimesToTop  int

	calX func() float64
}

func NewUFO(ld *loader.Loader) *UFO {
	entity := graphics.NewImageEntityWithImage(ld.ImageLoader.GetUFOImage(), ld.Cfg.ScreenWidth, ld.Cfg.ScreenHeight)

	entity.SetX(float64(rand.Intn(ld.Cfg.ScreenWidth-entity.Img.Width*2) + entity.Img.Width))
	entity.SetY(-float64(entity.Img.Height))

	u := &UFO{
		ImageEntity:  *entity,
		YSpeedFactor: 3,
		XSpeedFactor: 2,
		XTimesToTop:  3 * 60,
	}

	maxTrail := utils.RandomForMinFloat64(entity.X, float64(ld.Cfg.ScreenWidth)-entity.X)
	if maxTrail > 100 {
		u.calX = brick.GenerateRandomStableTrail(entity.X, 0, float64(ld.Cfg.ScreenHeight), u.XSpeedFactor)
	} else {
		u.calX = brick.GenerateRandomSinTrailWithBase(entity.X, maxTrail, u.XTimesToTop)
	}

	return u
}

/*
Update: update the UFO position (minus speedFactor).
return whether the UFO is out of the screen (for removing it from the storage).
*/
func (u *UFO) Update() bool {
	if u.IsDownOfScreen() {
		return false
	}

	u.MoveDown(u.YSpeedFactor)
	u.SetX(u.calX())
	return true
}
