package listener

import "github.com/zq-xu/go-game/internal/ebiten_game/gamedata"

type GameDataListener interface {
	gamedata.GameData
}

type gameDataListener struct {
	gamedata.GameData
}

func NewGameDataListener() GameDataListener {
	return &gameDataListener{
		GameData: gamedata.NewGameData(),
	}
}
