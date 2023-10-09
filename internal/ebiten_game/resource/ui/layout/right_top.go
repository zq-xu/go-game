package layout

import "github.com/ebitenui/ebitenui/widget"

type RightTopLayoutResource struct{}

func NewRightTopLayoutResource() *RightTopLayoutResource {
	return &RightTopLayoutResource{}
}

func (cr *RightTopLayoutResource) NewRightTopRowLayout(space, paddingRight int, fn func(row *widget.Container)) *widget.Container {
	root := newAnchorContainer()
	c := cr.newRowLayout(space, paddingRight)
	root.AddChild(c)

	fn(c)

	return root
}

func (cr *RightTopLayoutResource) newRowLayout(space, paddingRight int) *widget.Container {
	return widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Spacing(space),
			widget.RowLayoutOpts.Padding(widget.Insets{
				Top:   10,
				Right: paddingRight,
			}),
		)),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionEnd,
				VerticalPosition:   widget.AnchorLayoutPositionStart,
			}),
		),
	)
}
