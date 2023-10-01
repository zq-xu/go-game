package config

import "image/color"

var Cfg = &Config{
	ScreenConfig: ScreenConfig{
		ScreenWidth:  960,
		ScreenHeight: 720,
	},

	Title: "Ebiten Game",

	TitleFontSize: 24,
	FontSize:      12,
	SmallFontSize: 12,

	BgColor: color.RGBA{0xff, 0xff, 0xff, 0xff},
}

type Config struct {
	ScreenConfig `json:",inline"`

	Title string `json:"title"`

	TitleFontSize int `json:"titleFontSize"`
	FontSize      int `json:"fontSize"`
	SmallFontSize int `json:"smallFontSize"`

	BgColor color.RGBA `json:"bgColor"`
}

type ScreenConfig struct {
	ScreenWidth  int `json:"screenWidth"`
	ScreenHeight int `json:"screenHeight"`
}
