package ui

import (
	"github.com/ebitenui/ebitenui/widget"
	"golang.org/x/image/font"

	"github.com/zq-xu/2d-game/internal/ebiten_game/resource/loader"
)

const (
	ButtonIdlePath         = "images/button/button-idle.png"
	ButtonHoverPath        = "images/button/button-hover.png"
	ButtonPressedHoverPath = "images/button/button-selected-hover.png"
	ButtonPressedPath      = "images/button/button-pressed.png"
	ButtonDisabledPath     = "images/button/button-disabled.png"
)

type ButtonResource struct {
	Image   *widget.ButtonImage
	Text    *widget.ButtonTextColor
	Face    font.Face
	Padding widget.Insets
}

func NewButtonImage(ld *loader.Loader) *widget.ButtonImage {
	idle := ld.ImageLoader.NineSliceImgLoader.MustGetNineSliceImage(ButtonIdlePath, 12, 0)
	hover := ld.ImageLoader.NineSliceImgLoader.MustGetNineSliceImage(ButtonHoverPath, 12, 0)
	pressed_hover := ld.ImageLoader.NineSliceImgLoader.MustGetNineSliceImage(ButtonPressedHoverPath, 12, 0)
	pressed := ld.ImageLoader.NineSliceImgLoader.MustGetNineSliceImage(ButtonPressedPath, 12, 0)
	disabled := ld.ImageLoader.NineSliceImgLoader.MustGetNineSliceImage(ButtonDisabledPath, 12, 0)

	return &widget.ButtonImage{
		Idle:         idle,
		Hover:        hover,
		Pressed:      pressed,
		PressedHover: pressed_hover,
		Disabled:     disabled,
	}
}

func NewButtonResource(ld *loader.Loader, cr *ColorResource) *ButtonResource {
	i := NewButtonImage(ld)

	return &ButtonResource{
		Image: i,

		Text: &widget.ButtonTextColor{
			Idle:     cr.TextIdleColor,
			Disabled: cr.TextDisabledColor,
		},

		Face: ld.FontLoader.Face(),

		Padding: widget.Insets{
			Left:   30,
			Right:  30,
			Top:    5,
			Bottom: 5,
		},
	}
}

func (b *ButtonResource) NewButton(text string, opts ...widget.ButtonOpt) *widget.Button {
	opts = append(opts,
		widget.ButtonOpts.Image(b.Image),
		widget.ButtonOpts.Text(text, b.Face, b.Text),
		widget.ButtonOpts.TextPadding(b.Padding),
		widget.ButtonOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
			Stretch: true,
		})),
	)

	return widget.NewButton(opts...)
}
