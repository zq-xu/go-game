package entity

import (
	"math/rand"

	"github.com/zq-xu/2d-game/internal/ebiten_game/game"
	"github.com/zq-xu/2d-game/internal/ebiten_game/resource/loader"
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

func NewUFO(ctx *game.Context) *UFO {
	img := ctx.Resource.ImageLoader.ImgLoader.MustGetImage(loader.UFOImgPath)
	entity := graphics.NewImageEntityWithImage(img, ctx.Resource.Cfg.ScreenWidth, ctx.Resource.Cfg.ScreenHeight)

	entity.UnlimitTop()
	entity.UnlimitBottom()

	entity.SetX(float64(rand.Intn(ctx.Resource.Cfg.ScreenWidth-entity.Img.Width*2) + entity.Img.Width))
	entity.SetY(-float64(entity.Img.Height))

	u := &UFO{
		ImageEntity:  *entity,
		YSpeedFactor: 3,
		XSpeedFactor: 2,
		XTimesToTop:  3 * 60,
	}

	maxTrail := utils.RandomForMinFloat64(entity.X, float64(ctx.Resource.Cfg.ScreenWidth)-entity.X)
	if maxTrail > 100 {
		u.calX = brick.GenerateRandomStableTrail(entity.X, 0, float64(ctx.Resource.Cfg.ScreenHeight), u.XSpeedFactor)
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
