package metric

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/zq-xu/go-game/pkg/graphics"
)

const (
	defaultStartX = 4

	MetricCharWidth  = 6
	MetricLineHeight = 16
)

var DefaultDrawConfig = &DrawConfig{
	Face:  *graphics.GetFont().Face(),
	X:     defaultStartX,
	Y:     MetricLineHeight,
	Color: color.Black,
}

type DrawConfig struct {
	Face  text.Face
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
