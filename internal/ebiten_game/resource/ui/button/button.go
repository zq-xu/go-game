package button

import (
	"github.com/zq-xu/2d-game/internal/ebiten_game/resource/loader"
	"github.com/zq-xu/2d-game/internal/ebiten_game/resource/ui/color"
)

type ButtonResource struct {
	MenuButtonResource
	SettingButtonResource
}

func NewButtonResource(ld *loader.Loader, cr *color.ColorResource) *ButtonResource {
	return &ButtonResource{
		MenuButtonResource:    *NewMenuButtonResource(ld, cr),
		SettingButtonResource: *NewSettingButtonResource(ld),
	}
}
