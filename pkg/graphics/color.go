package graphics

import (
	"image/color"
	"strconv"
)

const (
	textIdleRGBA     = "dff4ff"
	textDisabledRGBA = "5a7a91"
)

var globalColor Color

type Color interface {
	TextIdleColor() color.Color
	TextDisabledColor() color.Color
}

type colorInstance struct {
	textIdleColor     color.Color
	textDisabledColor color.Color
}

func GetColor() Color {
	if globalColor == nil {
		globalColor = newColor()
	}
	return globalColor
}

func newColor() Color {
	return &colorInstance{
		textIdleColor:     HexToColor(textIdleRGBA),
		textDisabledColor: HexToColor(textDisabledRGBA),
	}
}

func (c *colorInstance) TextIdleColor() color.Color { return c.textIdleColor }

func (c *colorInstance) TextDisabledColor() color.Color { return c.textDisabledColor }

func HexToColor(h string) color.Color {
	u, err := strconv.ParseUint(h, 16, 0)
	if err != nil {
		panic(err)
	}

	return color.NRGBA{
		R: uint8(u & 0xff0000 >> 16),
		G: uint8(u & 0xff00 >> 8),
		B: uint8(u & 0xff),
		A: 255,
	}
}
