package config

import (
	"github.com/zq-xu/gotools/configx"
)

type Config struct{}

func init() {
	configx.RegisterByFile("tiledmap", &TiledMapCfg, configx.DefaultSetupFunc)
	configx.RegisterByFile("dialog", &DialogCfg, configx.DefaultSetupFunc)

}
