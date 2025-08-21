package button

import (
	"github.com/ebitenui/ebitenui/widget"
	"github.com/rotisserie/eris"
	"github.com/sirupsen/logrus"

	"github.com/zq-xu/go-game/pkg/graphics"
)

const (
	SettingButtonIdlePath         = "images/button/setting/setting-button-idle.png"
	SettingButtonHoverPath        = "images/button/setting/setting-button-hover.png"
	SettingButtonPressedHoverPath = "images/button/setting/setting-button-selected-hover.png"
	SettingButtonPressedPath      = "images/button/setting/setting-button-pressed.png"
	SettingButtonDisabledPath     = "images/button/setting/setting-button-disabled.png"
)

var settingButtonImage *widget.ButtonImage

// NewSettingButton
func NewSettingButton(opts ...widget.ButtonOpt) *widget.Button {
	settingButtonImage, err := getSettingButtonImage()
	if err != nil {
		logrus.Fatalf("get setting button image failed. %s", err)
		return nil
	}

	opts = append(opts,
		widget.ButtonOpts.Image(settingButtonImage),
		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{Stretch: true}),
		),
	)

	return widget.NewButton(opts...)
}

func getSettingButtonImage() (*widget.ButtonImage, error) {
	if settingButtonImage != nil {
		return settingButtonImage, nil
	}

	var err error
	bi := &widget.ButtonImage{}

	bi.Idle, err = graphics.GetNineSliceImage(SettingButtonIdlePath, 50, 42)
	if err != nil {
		return nil, eris.Wrap(err, "get setting idle image failed.")
	}

	bi.Hover, err = graphics.GetNineSliceImage(SettingButtonHoverPath, 50, 42)
	if err != nil {
		return nil, eris.Wrap(err, "get setting hover image failed.")
	}

	bi.PressedHover, err = graphics.GetNineSliceImage(SettingButtonPressedHoverPath, 50, 42)
	if err != nil {
		return nil, eris.Wrap(err, "get setting pressed hover image failed.")
	}

	bi.Pressed, err = graphics.GetNineSliceImage(SettingButtonPressedPath, 50, 42)
	if err != nil {
		return nil, eris.Wrap(err, "get setting pressed image failed.")
	}

	bi.Disabled, err = graphics.GetNineSliceImage(SettingButtonDisabledPath, 50, 42)
	if err != nil {
		return nil, eris.Wrap(err, "get setting disabled image failed.")
	}

	return bi, nil
}
