package gaming

import (
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/internal/shooter/data"
	"github.com/zq-xu/go-game/internal/shooter/ui/components"
	"github.com/zq-xu/go-game/internal/shooter/ui/layout"
)

const (
	metricWidth  = 300
	metricHeight = 300
)

type metrics struct {
	gameData data.Data

	img *ebiten.Image
	ui  *ebitenui.UI

	shadowDrawer func(screen *ebiten.Image)
}

func NewMetric(gameData data.Data) *metrics {
	g := &metrics{
		gameData:     gameData,
		img:          ebiten.NewImage(metricWidth, metricHeight),
		shadowDrawer: components.GenerateShadowDrawerFn(0),
	}

	g.initUI()
	return g
}

func (g *metrics) Update() error {
	g.ui.Update()
	return nil
}

func (g *metrics) Draw(screen *ebiten.Image) {
	g.img.Clear()
	g.shadowDrawer(g.img)
	g.gameData.Metrics().Draw(g.img)

	g.ui.Draw(screen)
}

func (g *metrics) initUI() {
	rootContainer := layout.NewLeftTopLayout(
		widget.ContainerOpts.BackgroundImage(
			image.NewNineSliceSimple(g.img, metricWidth, metricHeight),
		),
	)

	g.ui = &ebitenui.UI{Container: rootContainer}
}
