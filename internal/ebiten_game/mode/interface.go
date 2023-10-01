package mode

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/2d-game/internal/ebiten_game/data"
)

const (
	WaitingStartMode Mode = iota
	GamingMode
	GameOverMode
)

var modeGeneratorSet = make(map[Mode]gameModeHandlerGenerator)

type Mode int

type gameModeHandlerGenerator func(gd *data.GameData, ls ModeListener) GameModeHandler

type GameModeHandler interface {
	Update() error
	Draw(screen *ebiten.Image)
	Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int)
}

type ModeListener interface {
	SetMode(m Mode)
}

func Register(m Mode, fn gameModeHandlerGenerator) {
	modeGeneratorSet[m] = fn
}

func NewModeHandler(m Mode, gd *data.GameData, ls ModeListener) (GameModeHandler, error) {
	fn, ok := modeGeneratorSet[m]
	if ok {
		return fn(gd, ls), nil
	}

	return nil, fmt.Errorf("invalid mode %d", m)
}
