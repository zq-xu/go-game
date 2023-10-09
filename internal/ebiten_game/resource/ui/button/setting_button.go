package button

import (
	"github.com/ebitenui/ebitenui/widget"
	"golang.org/x/image/font"

	"github.com/zq-xu/2d-game/internal/ebiten_game/resource/loader"
)

const (
	SettingButtonIdlePath         = "images/button/setting/setting-button-idle.png"
	SettingButtonHoverPath        = "images/button/setting/setting-button-hover.png"
	SettingButtonPressedHoverPath = "images/button/setting/setting-button-selected-hover.png"
	SettingButtonPressedPath      = "images/button/setting/setting-button-pressed.png"
	SettingButtonDisabledPath     = "images/button/setting/setting-button-disabled.png"
)

type SettingButtonResource struct {
	Image *widget.ButtonImage

	Face    font.Face
	Padding widget.Insets
}

func (mb *SettingButtonResource) newButtonImage(ld *loader.Loader) *widget.ButtonImage {
	idle := ld.ImageLoader.NineSliceImgLoader.MustGetNineSliceSimpleImage(SettingButtonIdlePath, 50, 42)
	hover := ld.ImageLoader.NineSliceImgLoader.MustGetNineSliceSimpleImage(SettingButtonHoverPath, 50, 42)
	pressed_hover := ld.ImageLoader.NineSliceImgLoader.MustGetNineSliceSimpleImage(SettingButtonPressedHoverPath, 50, 42)
	pressed := ld.ImageLoader.NineSliceImgLoader.MustGetNineSliceSimpleImage(SettingButtonPressedPath, 50, 42)
	disabled := ld.ImageLoader.NineSliceImgLoader.MustGetNineSliceSimpleImage(SettingButtonDisabledPath, 50, 42)

	return &widget.ButtonImage{
		Idle:         idle,
		Hover:        hover,
		Pressed:      pressed,
		PressedHover: pressed_hover,
		Disabled:     disabled,
	}
}

func NewSettingButtonResource(ld *loader.Loader) *SettingButtonResource {
	mb := &SettingButtonResource{}

	mb.Image = mb.newButtonImage(ld)

	return mb
}

func (b *SettingButtonResource) NewSettingButton(opts ...widget.ButtonOpt) *widget.Button {
	opts = append(opts,
		widget.ButtonOpts.Image(b.Image),

		widget.ButtonOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
			Stretch: true,
		})),
	)

	return widget.NewButton(opts...)
}
