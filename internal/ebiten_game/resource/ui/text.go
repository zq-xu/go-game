package ui

import (
	"image/color"

	"github.com/ebitenui/ebitenui/widget"
	"golang.org/x/image/font"

	uiColor "github.com/zq-xu/go-game/internal/ebiten_game/resource/ui/color"
)

type TextResource struct {
	Color color.Color
}

func NewTextResource(cr *uiColor.ColorResource) *TextResource {
	return &TextResource{
		Color: cr.TextIdleColor,
	}
}

func (tr *TextResource) NewCenterText(txt string, face font.Face) *widget.Text {
	return widget.NewText(widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
		HorizontalPosition: widget.AnchorLayoutPositionCenter,
		VerticalPosition:   widget.AnchorLayoutPositionCenter,
	})),
		widget.TextOpts.Position(widget.TextPositionCenter, widget.TextPositionCenter),
		widget.TextOpts.Text(txt, face, tr.Color))
}
