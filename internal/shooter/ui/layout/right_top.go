package layout

import "github.com/ebitenui/ebitenui/widget"

func NewRightTopRowLayout(space, paddingRight int, fn func(row *widget.Container)) *widget.Container {
	c := newRowLayout(space, paddingRight)
	fn(c)

	root := newAnchorContainer()
	root.AddChild(c)
	return root
}

func newRowLayout(space, paddingRight int) *widget.Container {
	return widget.NewContainer(
		widget.ContainerOpts.Layout(
			widget.NewRowLayout(
				widget.RowLayoutOpts.Spacing(space),
				widget.RowLayoutOpts.Padding(&widget.Insets{
					Top:   10,
					Right: paddingRight,
				}),
			)),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(
				widget.AnchorLayoutData{
					HorizontalPosition: widget.AnchorLayoutPositionEnd,
					VerticalPosition:   widget.AnchorLayoutPositionStart,
				}),
		),
	)
}
