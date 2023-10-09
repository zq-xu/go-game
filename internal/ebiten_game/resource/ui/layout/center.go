package layout

import "github.com/ebitenui/ebitenui/widget"

type CenterLayoutResource struct{}

func NewCenterLayoutResource() *CenterLayoutResource {
	return &CenterLayoutResource{}
}

func (cr *CenterLayoutResource) NewCenterRowLayout(minWidth, spacing int, rowscale []bool, fn func(row *widget.Container)) *widget.Container {
	root := newAnchorContainer()
	c := newSingleColumnGridLayout(minWidth, spacing, rowscale)
	root.AddChild(c)

	fn(c)

	return root
}

func newAnchorContainer() *widget.Container {
	return widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			StretchHorizontal: true,
		})),
	)
}

func newSingleColumnGridLayout(minWidth, spacing int, rowscale []bool) *widget.Container {
	return widget.NewContainer(
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.MinSize(minWidth, 0)),
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			StretchHorizontal: true,
			StretchVertical:   true,
		})),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
			}),
		),
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(1),
			widget.GridLayoutOpts.Stretch([]bool{true}, rowscale),
			widget.GridLayoutOpts.Spacing(spacing, spacing),
		)),
	)
}
