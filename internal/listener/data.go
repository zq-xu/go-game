package listener

import "github.com/zq-xu/go-game/internal/data"

type GameDataListener interface {
	data.Data
}

type gameDataListener struct {
	data.Data
}

func NewGameDataListener() GameDataListener {
	return &gameDataListener{
		Data: data.NewGameData(),
	}
}
