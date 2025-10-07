package base

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/pkg/logs"
)

// ShowGifs
func ShowGifs(imgPaths ...string) {
	logs.InitLogger("debug")

	g, err := NewGame(imgPaths...)
	if err != nil {
		logs.Logger.Fatalf("new image from file %s failed. %v", imgPaths, err)
	}

	ebiten.SetWindowSize(g.ScreenWidth, g.ScreenHeight)
	ebiten.SetWindowTitle("Gif Test")

	err = ebiten.RunGame(g)
	if err != nil {
		logs.Logger.Fatal("run game failed.", err)
	}
}
