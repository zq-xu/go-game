package main

import (
	"log"

	"github.com/zq-xu/go-game/internal/dungeon"
	"github.com/zq-xu/go-game/pkg/logs"
)

func main() {
	logs.InitLogger("debug")

	err := dungeon.StartGame()
	if err != nil {
		log.Fatal(err)
	}
}
