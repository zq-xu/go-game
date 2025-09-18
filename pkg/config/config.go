package config

import (
	"image/color"
)

var Cfg = &Config{
	ScreenConfig: ScreenConfig{
		FullScreen:   true,
		ScreenWidth:  int(BaseScreenWidth),
		ScreenHeight: int(BaseScreenHeight),
	},

	Title:          "Ebiten Game",
	AuthorText:     "AUTHOR: ZHIQIANG XU",
	StartHintTexts: []string{"", "", "PRESS ENTER KEY TO START", "", "", "", ""},

	BgColor:     color.RGBA{0xff, 0xff, 0xff, 0xff},
	ShadowColor: color.RGBA{0x30, 0x30, 0x30, 0x30},

	KeyInterval: 500,
}

type Config struct {
	ScreenConfig ScreenConfig `json:",inline"`

	Title          string   `json:"title"`
	AuthorText     string   `json:"authorText"`
	StartHintTexts []string `json:"startHintTexts"`

	BgColor color.RGBA `json:"bgColor"`

	ShadowColor color.RGBA `json:"shadowColor"`

	KeyInterval int `json:"keyInterval" description:"base on the TPS"`
}

type ScreenConfig struct {
	FullScreen   bool `json:"fullScreen"`
	ScreenWidth  int  `json:"screenWidth"`
	ScreenHeight int  `json:"screenHeight"`
}
