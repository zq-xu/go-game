package npc

import (
	"github.com/zq-xu/gotools/configx"

	"github.com/zq-xu/go-game/internal/dungeon/entity/core/actor"
	"github.com/zq-xu/go-game/internal/dungeon/types"
)

var NPCConfigSet = make(map[string]*NPCConfig, 0)

type NPCConfig struct {
	actor.Option              `mapstructure:",squash"`
	DialogOffsetToRightCenter *types.Offset
}

func init() {
	configx.RegisterByFile("npcs", &NPCConfigSet, configx.DefaultSetupFunc)
}
