package main

import (
	"log"

	"github.com/zq-xu/go-game/internal/dungeon"
	"github.com/zq-xu/go-game/pkg/logs"
)

func main() {
	err := logs.InitLog("debug")
	if err != nil {
		log.Fatal(err)
	}

	err = dungeon.StartGame()
	if err != nil {
		log.Fatal(err)
	}
}
