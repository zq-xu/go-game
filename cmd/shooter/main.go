package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/internal/shooter/game"
	"github.com/zq-xu/go-game/pkg/logs"
)

func main() {
	err := logs.InitLog("debug")
	if err != nil {
		log.Fatal(err)
	}

	err = ebiten.RunGame(game.NewGame())
	if err != nil {
		log.Fatal(err)
	}
}
