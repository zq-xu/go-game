package base

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/samber/lo"

	"github.com/zq-xu/go-game/pkg/graphics/images/gifkit"
	"github.com/zq-xu/go-game/pkg/graphics/images/imagekit"
	"github.com/zq-xu/go-game/pkg/graphics/images/imagetable"
	"github.com/zq-xu/go-game/pkg/logs"
)

type game struct {
	gifs []gifkit.Gif

	maxWidth  int
	maxHeight int

	locations                 []image.Point
	ScreenWidth, ScreenHeight int
}

func (g *game) Update() error { return nil }

func (g *game) Draw(screen *ebiten.Image) {
	for k, v := range g.gifs {
		v.Draw(screen, float64(g.locations[k].X), float64(g.locations[k].Y))
	}
}

func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.ScreenWidth, g.ScreenHeight
}

func NewGame(imgPaths ...string) (*game, error) {
	g := &game{}
	for _, imgPath := range imgPaths {
		img, err := imagekit.NewImageFromFile(imgPath)
		if err != nil {
			logs.Logger.Fatalf("new image from file %s failed. %v", imgPath, err)
		}

		imgTable := imagetable.NewImageTable(imgPath, img)
		imgTable.LogBoxes()

		gif := gifkit.NewNeverStopGif(imgTable.Images()...)
		gif.SetDrawBeginning(gifkit.CenterDrawBeginning)
		gif.SetUpdateInterval(10)
		g.addGif(gif)
	}

	g.loadLocationsAndScreenSize()
	return g, nil
}

func (g *game) addGif(gif gifkit.Gif) {
	g.gifs = append(g.gifs, gif)
	g.maxWidth = lo.Max([]int{gif.Width(), g.maxWidth})
	g.maxHeight = lo.Max([]int{gif.Height(), g.maxHeight})
}

func (g *game) loadLocationsAndScreenSize() {
	g.ScreenWidth = (g.maxWidth+100)*len(g.gifs) + 100
	g.ScreenHeight = g.maxHeight + 200

	g.locations = make([]image.Point, len(g.gifs))
	wStart, h := 0, g.ScreenHeight/2
	for k, v := range g.gifs {
		wStart += 100
		g.locations[k] = image.Point{X: wStart + v.Width()/2, Y: h}
		wStart += v.Width()
	}
}
