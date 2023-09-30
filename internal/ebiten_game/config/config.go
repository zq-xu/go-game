package config

import "image/color"

var Cfg = &Config{
	ScreenConfig: ScreenConfig{
		ScreenWidth:  960,
		ScreenHeight: 720,
	},

	Title:         "Ebiten Game",
	BgColor:       color.RGBA{0xff, 0xff, 0xff, 0xff},
	ImageRootPath: "../../assets/image",
}

type Config struct {
	ScreenConfig `json:",inline"`

	Title         string     `json:"title"`
	BgColor       color.RGBA `json:"bgColor"`
	ImageRootPath string     `json:"imageRootPath"`
}

type ScreenConfig struct {
	ScreenWidth  int `json:"screenWidth"`
	ScreenHeight int `json:"screenHeight"`
}
