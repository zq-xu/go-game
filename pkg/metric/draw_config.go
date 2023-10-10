package metric

import (
	"image/color"

	"github.com/hajimehoshi/bitmapfont/v3"
	"golang.org/x/image/font"
)

const (
	defaultStartX = 4

	MetricCharWidth  = 6
	MetricLineHeight = 16
)

var DefaultDrawConfig = &DrawConfig{
	Face:  bitmapfont.Face,
	X:     defaultStartX,
	Y:     MetricLineHeight,
	Color: color.Black,
}

type DrawConfig struct {
	Face  font.Face
	X     int
	Y     int
	Color color.Color
}

func NewDrawConfig() *DrawConfig {
	return &DrawConfig{}
}

func (cfg *DrawConfig) Copy() *DrawConfig {
	return &DrawConfig{
		Face:  cfg.Face,
		X:     cfg.X,
		Y:     cfg.Y,
		Color: cfg.Color,
	}
}
