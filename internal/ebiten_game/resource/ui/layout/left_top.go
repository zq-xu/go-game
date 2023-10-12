package layout

import (
	"github.com/ebitenui/ebitenui/widget"
)

type LeftTopLayoutResource struct{}

func NewLeftTopLayoutResource() *LeftTopLayoutResource {
	return &LeftTopLayoutResource{}
}

func (cr *LeftTopLayoutResource) NewLeftTopLayout(opts ...widget.ContainerOpt) *widget.Container {
	opts = append(opts,
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.MinSize(500, 500)),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				StretchHorizontal: true,
			}),
		),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionStart,
				VerticalPosition:   widget.AnchorLayoutPositionStart,
			}),
		))

	return widget.NewContainer(opts...)
}
