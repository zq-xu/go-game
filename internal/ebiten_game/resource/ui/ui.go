package ui

import (
	"github.com/zq-xu/2d-game/internal/ebiten_game/config"
	"github.com/zq-xu/2d-game/internal/ebiten_game/resource/loader"
	"github.com/zq-xu/2d-game/internal/ebiten_game/resource/ui/button"
	"github.com/zq-xu/2d-game/internal/ebiten_game/resource/ui/color"
	uiColor "github.com/zq-xu/2d-game/internal/ebiten_game/resource/ui/color"
	"github.com/zq-xu/2d-game/internal/ebiten_game/resource/ui/layout"
)

type UIResource struct {
	ColorResource *uiColor.ColorResource

	BackgroundResource *BackgroundResource
	ShadowResource     *ShadowResource

	SeparatorResource *SeparatorResource
	ButtonResource    *button.ButtonResource
	TextResource      *TextResource

	LayoutResource *layout.LayoutResource
}

func NewUIResource(cfg *config.Config, ld *loader.Loader) *UIResource {
	ur := &UIResource{}

	ur.ColorResource = color.NewColorResource()
	ur.BackgroundResource = NewBackgroundResource()
	ur.ShadowResource = NewShadowResource(cfg)

	ur.SeparatorResource = NewSeparatorResource()
	ur.ButtonResource = button.NewButtonResource(ld, ur.ColorResource)
	ur.TextResource = NewTextResource(ur.ColorResource)

	ur.LayoutResource = layout.NewLayoutResource()

	return ur
}
