package layout

import "github.com/ebitenui/ebitenui/widget"

func newAnchorContainer() *widget.Container {
	return widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)
}
