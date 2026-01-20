package base

import (
	"image"

	"github.com/zq-xu/gotools/configx"

	"github.com/zq-xu/go-game/internal/dungeon/entity/core/actor"
)

var MonsterConfigSet = make(map[string]MonsterConfig, 0)

type MonsterConfig struct {
	actor.Option `mapstructure:",squash"`
	ActiveRange  *image.Rectangle
}

func init() {
	configx.RegisterByFile("monsters", &MonsterConfigSet, configx.DefaultSetupFunc)
}
