package base

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zq-xu/gotools/logx"
)

// ShowGifs
func ShowGifs(imgPaths ...string) {
	logx.InitLogger()

	g, err := NewGame(imgPaths...)
	if err != nil {
		logx.Logger.Fatalf("new image from file %s failed. %v", imgPaths, err)
	}

	ebiten.SetWindowSize(g.ScreenWidth, g.ScreenHeight)
	ebiten.SetWindowTitle("Gif Test")

	err = ebiten.RunGame(g)
	if err != nil {
		logx.Logger.Fatal("run game failed.", err)
	}
}
