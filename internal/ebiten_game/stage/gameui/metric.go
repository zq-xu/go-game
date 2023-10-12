package gameui

import (
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/2d-game/internal/ebiten_game/game"
	"github.com/zq-xu/2d-game/pkg/metric"
)

const (
	metricWidth  = 200
	metricHeight = 200
)

type Metric struct {
	ctx *game.Context

	img *ebiten.Image

	ui *ebitenui.UI

	shadowDrawer func(screen *ebiten.Image)
}

func NewMetric(ctx *game.Context) *Metric {
	g := &Metric{ctx: ctx}

	g.img = ebiten.NewImage(metricWidth, metricHeight)
	g.shadowDrawer = g.ctx.Resource.ShadowResource.GenerateShadowDrawerFn(0)

	g.ui = g.newUI()
	return g
}

func (g *Metric) Update() error {
	g.ui.Update()
	return nil
}

func (g *Metric) Draw(screen *ebiten.Image) {
	g.ui.Draw(screen)

	g.img.Clear()
	g.shadowDrawer(g.img)
	metric.MultiPool.Draw(g.img)
}

func (g *Metric) newUI() *ebitenui.UI {
	rootContainer := g.ctx.Resource.LayoutResource.NewLeftTopLayout(
		widget.ContainerOpts.BackgroundImage(
			image.NewNineSliceSimple(g.img, metricWidth, metricHeight),
		),
	)

	return &ebitenui.UI{Container: rootContainer}
}
