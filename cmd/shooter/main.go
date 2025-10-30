package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zq-xu/gotools/configx"

	"github.com/zq-xu/go-game/internal/shooter/game"
)

var configFile = "etc/shooter/config.yaml"

func main() {
	err := configx.Setup(configFile)
	if err != nil {
		log.Fatal(err)
	}

	g, err := game.NewGame()
	if err != nil {
		log.Fatal(err)
	}

	err = ebiten.RunGame(g)
	if err != nil {
		log.Fatal(err)
	}
}
