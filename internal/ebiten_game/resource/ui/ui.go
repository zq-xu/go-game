package ui

import (
	"github.com/zq-xu/2d-game/internal/ebiten_game/resource/loader"
)

type UIResource struct {
	ColorResource      *ColorResource
	BackgroundResource *BackgroundResource

	SeparatorResource *SeparatorResource
	ButtonResource    *ButtonResource
	TextResource      *TextResource

	LayoutResource *LayoutResource
}

func NewUIResource(ld *loader.Loader) *UIResource {
	ur := &UIResource{}

	ur.ColorResource = NewColorResource()
	ur.BackgroundResource = NewBackgroundResource()

	ur.SeparatorResource = NewSeparatorResource()
	ur.ButtonResource = NewButtonResource(ld, ur.ColorResource)
	ur.TextResource = NewTextResource(ur.ColorResource)

	ur.LayoutResource = NewLayoutResource()

	return ur
}
