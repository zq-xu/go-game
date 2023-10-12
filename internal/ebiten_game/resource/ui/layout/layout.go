package layout

import "github.com/ebitenui/ebitenui/widget"

type LayoutResource struct {
	CenterLayoutResource
	RightTopLayoutResource
	LeftTopLayoutResource
}

func NewLayoutResource() *LayoutResource {
	return &LayoutResource{
		CenterLayoutResource:   *NewCenterLayoutResource(),
		RightTopLayoutResource: *NewRightTopLayoutResource(),
		LeftTopLayoutResource:  *NewLeftTopLayoutResource(),
	}
}

func newAnchorContainer() *widget.Container {
	return widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			StretchHorizontal: true,
		})),
	)
}
