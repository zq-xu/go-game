package config

import (
	"github.com/zq-xu/gotools/configx"

	"github.com/zq-xu/go-game/internal/dungeon/entity/core/tiledmap"
)

var TiledMapCfg tiledmap.TiledMapConfig

func GetTiledmapConfig() *tiledmap.TiledMapConfig {
	return &TiledMapCfg
}

func init() {
	configx.RegisterByFile("tiledmap", &TiledMapCfg, configx.DefaultSetupFunc)
}
