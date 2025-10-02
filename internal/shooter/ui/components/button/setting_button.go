package button

import (
	"github.com/ebitenui/ebitenui/widget"
	"github.com/rotisserie/eris"
	"github.com/sirupsen/logrus"

	"github.com/zq-xu/go-game/assets"
	"github.com/zq-xu/go-game/internal/shooter/ui/resources"
)

var (
	SettingButtonIdlePath         = assets.GetShooterImagePath("button/setting/setting-button-idle.png")
	SettingButtonHoverPath        = assets.GetShooterImagePath("button/setting/setting-button-hover.png")
	SettingButtonPressedHoverPath = assets.GetShooterImagePath("button/setting/setting-button-selected-hover.png")
	SettingButtonPressedPath      = assets.GetShooterImagePath("button/setting/setting-button-pressed.png")
	SettingButtonDisabledPath     = assets.GetShooterImagePath("button/setting/setting-button-disabled.png")
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

	bi.Idle, err = resources.GetFixedNineSlice(SettingButtonIdlePath)
	if err != nil {
		return nil, eris.Wrap(err, "get setting idle image failed.")
	}

	bi.Hover, err = resources.GetFixedNineSlice(SettingButtonHoverPath)
	if err != nil {
		return nil, eris.Wrap(err, "get setting hover image failed.")
	}

	bi.PressedHover, err = resources.GetFixedNineSlice(SettingButtonPressedHoverPath)
	if err != nil {
		return nil, eris.Wrap(err, "get setting pressed hover image failed.")
	}

	bi.Pressed, err = resources.GetFixedNineSlice(SettingButtonPressedPath)
	if err != nil {
		return nil, eris.Wrap(err, "get setting pressed image failed.")
	}

	bi.Disabled, err = resources.GetFixedNineSlice(SettingButtonDisabledPath)
	if err != nil {
		return nil, eris.Wrap(err, "get setting disabled image failed.")
	}

	return bi, nil
}
