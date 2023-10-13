package ui

import (
	"image/color"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"

	"github.com/zq-xu/go-game/pkg/graphics"
)

const separatorRGBA = "2a3944"

type SeparatorResource struct {
	Color color.Color
}

func NewSeparatorResource() *SeparatorResource {
	return &SeparatorResource{
		Color: graphics.HexToColor(separatorRGBA),
	}
}

func (sr *SeparatorResource) NewSeparator(it interface{}) widget.PreferredSizeLocateableWidget {
	c := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Padding(widget.Insets{
				Top:    20,
				Bottom: 20,
			}))),
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(it)))

	c.AddChild(widget.NewGraphic(
		widget.GraphicOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
			Stretch:   true,
			MaxHeight: 2,
		})),
		widget.GraphicOpts.ImageNineSlice(image.NewNineSliceColor(sr.Color)),
	))

	return c
}
