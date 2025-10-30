package main

import (
	"log"

	"github.com/zq-xu/gotools/configx"

	"github.com/zq-xu/go-game/internal/dungeon"
)

var configFile = "etc/dungeon/config.yaml"

func main() {
	err := configx.Setup(configFile)
	if err != nil {
		log.Fatal(err)
	}

	err = dungeon.StartGame()
	if err != nil {
		log.Fatal(err)
	}
}
