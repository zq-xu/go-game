package data

import "time"

type Data interface {
	AddShotUFO(count int)
	GetShotUFO() int64
}

type gameData struct {
	startTime    time.Time
	endTime      time.Time
	timeDuration time.Duration

	shotUFO int64
	score   int64
}

func NewGameData() Data {
	return &gameData{}
}

func (gd *gameData) AddShotUFO(count int) {
	gd.shotUFO += int64(count)
}
func (gd *gameData) GetShotUFO() int64 {
	return gd.shotUFO
}
