package config

import "image/color"

var Cfg = &Config{
	ScreenConfig: ScreenConfig{
		FullScreen:   true,
		ScreenWidth:  1920,
		ScreenHeight: 1080,
	},

	Title: "Ebiten Game",

	TitleFontSize: 24,
	FontSize:      12,
	SmallFontSize: 12,

	BgColor: color.RGBA{0xff, 0xff, 0xff, 0xff},

	ShipXSpeedFactor:  10,
	ShipYSpeedFactor:  5,
	BulletSpeedFactor: 10,
}

type Config struct {
	ScreenConfig `json:",inline"`

	Title string `json:"title"`

	TitleFontSize int `json:"titleFontSize"`
	FontSize      int `json:"fontSize"`
	SmallFontSize int `json:"smallFontSize"`

	BgColor color.RGBA `json:"bgColor"`

	ShipXSpeedFactor  float64 `json:"shipXSpeedFactor"`
	ShipYSpeedFactor  float64 `json:"shipYSpeedFactor"`
	BulletSpeedFactor float64 `json:"bulletSpeedFactor"`
}

type ScreenConfig struct {
	FullScreen   bool `json:"fullScreen"`
	ScreenWidth  int  `json:"screenWidth"`
	ScreenHeight int  `json:"screenHeight"`
}

func (c *ScreenConfig) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return c.ScreenWidth, c.ScreenHeight
}
