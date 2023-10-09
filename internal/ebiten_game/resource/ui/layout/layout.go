package layout

type LayoutResource struct {
	CenterLayoutResource
	RightTopLayoutResource
}

func NewLayoutResource() *LayoutResource {
	return &LayoutResource{
		CenterLayoutResource:   *NewCenterLayoutResource(),
		RightTopLayoutResource: *NewRightTopLayoutResource(),
	}
}
