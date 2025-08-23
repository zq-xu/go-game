package text

import (
	"image/color"

	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2/text/v2"

	"github.com/zq-xu/go-game/pkg/graphics"
)

func NewCenterText(txt string, f graphics.Font, cr color.Color) *widget.Text {
	return newCenterText(txt, f.Face(), cr)
}

func NewCenterBoldText(txt string, f graphics.Font, cr color.Color) *widget.Text {
	return newCenterText(txt, f.BoldFace(), cr)
}

func newCenterText(txt string, f *text.Face, cr color.Color) *widget.Text {
	return widget.NewText(
		widget.TextOpts.Text(txt, f, cr),
		widget.TextOpts.Position(widget.TextPositionCenter, widget.TextPositionCenter),
		widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
			}),
		))
}
