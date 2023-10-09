package config

import "image/color"

var Cfg = &Config{
	ScreenConfig: ScreenConfig{
		FullScreen:   true,
		ScreenWidth:  int(BaseScreenWidth),
		ScreenHeight: int(BaseScreenHeight),
	},

	Title:          "Ebiten Game",
	AuthorText:     "AUTHOR: ZHIQIANG XU",
	StartHintTexts: []string{"", "", "PRESS ENTER KEY TO START", "", "", "", ""},

	BgColor: color.RGBA{0xff, 0xff, 0xff, 0xff},
}

type Config struct {
	ScreenConfig ScreenConfig `json:",inline"`

	Title          string   `json:"title"`
	AuthorText     string   `json:"authorText"`
	StartHintTexts []string `json:"startHintTexts"`

	BgColor color.RGBA `json:"bgColor"`
}

type ScreenConfig struct {
	FullScreen   bool `json:"fullScreen"`
	ScreenWidth  int  `json:"screenWidth"`
	ScreenHeight int  `json:"screenHeight"`
}
