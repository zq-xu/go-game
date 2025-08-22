package gaming

import (
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/internal/ui/components"
	"github.com/zq-xu/go-game/internal/ui/layout"
	"github.com/zq-xu/go-game/pkg/metric"
)

const (
	metricWidth  = 200
	metricHeight = 200
)

type Metric struct {
	img *ebiten.Image

	ui *ebitenui.UI

	shadowDrawer func(screen *ebiten.Image)
}

func NewMetric() *Metric {
	g := &Metric{}

	g.img = ebiten.NewImage(metricWidth, metricHeight)
	g.shadowDrawer = components.GenerateShadowDrawerFn(0)

	g.ui = g.newUI()
	return g
}

func (g *Metric) Update() error {
	g.ui.Update()
	return nil
}

func (g *Metric) Draw(screen *ebiten.Image) {
	g.img.Clear()
	g.shadowDrawer(g.img)
	metric.MultiPool.Draw(g.img)

	g.ui.Draw(screen)
}

func (g *Metric) newUI() *ebitenui.UI {
	rootContainer := layout.NewLeftTopLayout(
		widget.ContainerOpts.BackgroundImage(
			image.NewNineSliceSimple(g.img, metricWidth, metricHeight),
		),
	)

	return &ebitenui.UI{Container: rootContainer}
}
