package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	ebitengame "github.com/zq-xu/2d-game/internal/ebiten_game"
)

//go:generate go install github.com/hajimehoshi/file2byteslice
//go:generate rm -rf ../../internal/ebiten_game/resources
//go:generate mkdir ../../internal/ebiten_game/resources
//go:generate file2byteslice -input ../../assets/image/ship.png -output ../../internal/ebiten_game/resources/ship.go -package resources -var ShipPng
//go:generate file2byteslice -input ../../assets/image/bullet.png -output ../../internal/ebiten_game/resources/bullet.go -package resources -var BulletPng
//go:generate file2byteslice -input ../../assets/image/ufo.png -output ../../internal/ebiten_game/resources/ufo.go -package resources -var UFOPng

func main() {

	// ebiten.SetWindowSize(640, 480)
	// ebiten.SetWindowTitle("Geometry Matrix")
	if err := ebiten.RunGame(ebitengame.NewGame()); err != nil {
		log.Fatal(err)
	}
}
