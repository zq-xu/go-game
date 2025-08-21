package components

import (
	"image/color"

	"github.com/ebitenui/ebitenui/widget"

	"github.com/zq-xu/go-game/pkg/graphics"
)

func NewCenterText(txt string, f graphics.Font, cr color.Color) *widget.Text {
	return widget.NewText(
		widget.TextOpts.Text(txt, f.Face(), cr),
		widget.TextOpts.Position(widget.TextPositionCenter, widget.TextPositionCenter),
		widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
			}),
		))
}
