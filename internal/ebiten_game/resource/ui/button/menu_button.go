package button

import (
	"github.com/ebitenui/ebitenui/widget"
	"golang.org/x/image/font"

	"github.com/zq-xu/2d-game/internal/ebiten_game/resource/loader"
	"github.com/zq-xu/2d-game/internal/ebiten_game/resource/ui/color"
)

const (
	MenuButtonIdlePath         = "images/button/menu/menu-button-idle.png"
	MenuButtonHoverPath        = "images/button/menu/menu-button-hover.png"
	MenuButtonPressedHoverPath = "images/button/menu/menu-button-selected-hover.png"
	MenuButtonPressedPath      = "images/button/menu/menu-button-pressed.png"
	MenuButtonDisabledPath     = "images/button/menu/menu-button-disabled.png"
)

type MenuButtonResource struct {
	Image     *widget.ButtonImage
	TextColor *widget.ButtonTextColor
	Face      font.Face
	Padding   widget.Insets
}

func (mb *MenuButtonResource) newButtonImage(ld *loader.Loader) *widget.ButtonImage {
	idle := ld.ImageLoader.NineSliceImgLoader.MustGetNineSliceImage(MenuButtonIdlePath, 12, 0)
	hover := ld.ImageLoader.NineSliceImgLoader.MustGetNineSliceImage(MenuButtonHoverPath, 12, 0)
	pressed_hover := ld.ImageLoader.NineSliceImgLoader.MustGetNineSliceImage(MenuButtonPressedHoverPath, 12, 0)
	pressed := ld.ImageLoader.NineSliceImgLoader.MustGetNineSliceImage(MenuButtonPressedPath, 12, 0)
	disabled := ld.ImageLoader.NineSliceImgLoader.MustGetNineSliceImage(MenuButtonDisabledPath, 12, 0)

	return &widget.ButtonImage{
		Idle:         idle,
		Hover:        hover,
		Pressed:      pressed,
		PressedHover: pressed_hover,
		Disabled:     disabled,
	}
}

func NewMenuButtonResource(ld *loader.Loader, cr *color.ColorResource) *MenuButtonResource {
	mb := &MenuButtonResource{}

	mb.Image = mb.newButtonImage(ld)

	mb.Face = ld.FontLoader.Face()

	mb.TextColor = &widget.ButtonTextColor{
		Idle:     cr.TextIdleColor,
		Disabled: cr.TextDisabledColor,
	}

	mb.Padding = widget.Insets{
		Left:   30,
		Right:  30,
		Top:    5,
		Bottom: 5,
	}

	return mb
}

func (b *MenuButtonResource) NewMenuButton(text string, opts ...widget.ButtonOpt) *widget.Button {
	opts = append(opts,
		widget.ButtonOpts.Image(b.Image),
		widget.ButtonOpts.Text(text, b.Face, b.TextColor),
		widget.ButtonOpts.TextPadding(b.Padding),
		widget.ButtonOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
			Stretch: true,
		})),
	)

	return widget.NewButton(opts...)
}
