package main

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/pkg/graphics/images/gifkit"
	"github.com/zq-xu/go-game/pkg/graphics/images/imagekit"
	"github.com/zq-xu/go-game/pkg/graphics/images/imagetable"
	"github.com/zq-xu/go-game/pkg/logs"
)

const (
	screenWidth  = 400
	screenHeight = 400

	imgPath = "blue_file_gif.png"
)

type game struct {
	gifkit.Gif
}

func (g *game) Update() error { return nil }

func (g *game) Draw(screen *ebiten.Image) {
	g.Gif.Draw(screen, screenWidth/2, screenHeight/2)
}

func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	logs.InitLogger("debug")

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Gif Test")

	img, err := imagekit.NewImageFromFile(imgPath)
	if err != nil {
		logs.Logger.Fatalf("new image from file %s failed. %v", imgPath, err)
	}

	imgTable := imagetable.NewImageTable("blue fire", img)
	imgTable.LogBoxes()

	g := gifkit.NewNeverStopGif(imgTable.Images()...)
	g.SetDrawBeginning(gifkit.CenterDrawBeginning)
	g.SetUpdateInterval(3)

	err = ebiten.RunGame(&game{Gif: g})
	if err != nil {
		logs.Logger.Fatal("run game failed.", err)
	}
}
