package button

import (
	"github.com/ebitenui/ebitenui/widget"
	"github.com/rotisserie/eris"
	"github.com/sirupsen/logrus"

	"github.com/zq-xu/go-game/pkg/graphics"
)

const (
	MenuButtonIdlePath         = "images/button/menu/menu-button-idle.png"
	MenuButtonHoverPath        = "images/button/menu/menu-button-hover.png"
	MenuButtonPressedHoverPath = "images/button/menu/menu-button-selected-hover.png"
	MenuButtonPressedPath      = "images/button/menu/menu-button-pressed.png"
	MenuButtonDisabledPath     = "images/button/menu/menu-button-disabled.png"
)

var menuButtonImage *widget.ButtonImage

// NewMenuButton
func NewMenuButton(text string, opts ...widget.ButtonOpt) *widget.Button {
	buttonImage, err := getMenuButtonImage()
	if err != nil {
		logrus.Fatalf("get menu button image failed. %s", err)
		return nil
	}

	opts = append(opts,
		widget.ButtonOpts.Image(buttonImage),
		widget.ButtonOpts.Text(text, graphics.GetFont().Face(), &widget.ButtonTextColor{
			Idle:     graphics.GetColor().TextIdleColor(),
			Disabled: graphics.GetColor().TextDisabledColor(),
		}),
		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{Stretch: true}),
		),
		widget.ButtonOpts.TextPadding(&widget.Insets{
			Left:   30,
			Right:  30,
			Top:    5,
			Bottom: 5,
		}),
	)

	return widget.NewButton(opts...)
}

func getMenuButtonImage() (*widget.ButtonImage, error) {
	if menuButtonImage != nil {
		return menuButtonImage, nil
	}

	var err error
	bi := &widget.ButtonImage{}

	bi.Idle, err = graphics.GetNineSliceImage(MenuButtonIdlePath, 12, 0)
	if err != nil {
		return nil, eris.Wrap(err, "get menu idle image failed.")
	}

	bi.Hover, err = graphics.GetNineSliceImage(MenuButtonHoverPath, 12, 0)
	if err != nil {
		return nil, eris.Wrap(err, "get menu hover image failed.")
	}

	bi.PressedHover, err = graphics.GetNineSliceImage(MenuButtonPressedHoverPath, 12, 0)
	if err != nil {
		return nil, eris.Wrap(err, "get menu pressed hover image failed.")
	}

	bi.Pressed, err = graphics.GetNineSliceImage(MenuButtonPressedPath, 12, 0)
	if err != nil {
		return nil, eris.Wrap(err, "get menu pressed image failed.")
	}

	bi.Disabled, err = graphics.GetNineSliceImage(MenuButtonDisabledPath, 12, 0)
	if err != nil {
		return nil, eris.Wrap(err, "get menu disabled image failed.")
	}

	return bi, nil
}
