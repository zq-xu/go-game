package config

import (
	"image/color"

	"github.com/zq-xu/gotools/configx"
)

var Cfg Config

func init() {
	configx.RegisterByFile("settings", &Cfg, configx.DefaultSetupFunc)
}

type Config struct {
	ScreenConfig ScreenConfig

	Title          string
	AuthorText     string
	StartHintTexts []string

	BgColor color.RGBA

	ShadowColor color.RGBA

	KeyInterval int `description:"base on the TPS"`
}

type ScreenConfig struct {
	FullScreen   bool
	ScreenWidth  int
	ScreenHeight int
}
