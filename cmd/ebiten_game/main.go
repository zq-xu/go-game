package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	ebitengame "github.com/zq-xu/go-game/internal/ebiten_game"
)

func main() {
	if err := ebiten.RunGame(ebitengame.NewGame()); err != nil {
		log.Fatal(err)
	}
}
