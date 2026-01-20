package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rotisserie/eris"
	"github.com/zq-xu/gotools/configx"

	"github.com/zq-xu/go-game/internal/dungeon"
	"github.com/zq-xu/go-game/internal/dungeon/config"
)

var configFile = "etc/dungeon/config.yaml"

func main() {
	err := configx.Setup(configFile)
	if err != nil {
		log.Fatal(err)
	}

	err = runDungeon()
	if err != nil {
		log.Fatal(err)
	}
}

func runDungeon() error {
	ebiten.SetWindowSize(config.ScreenWidth, config.ScreenHeight)
	ebiten.SetWindowTitle("Dungeon")

	g, err := dungeon.NewGame()
	if err != nil {
		return eris.Wrap(err, "failed to initialize game")
	}

	return ebiten.RunGame(g)
}
