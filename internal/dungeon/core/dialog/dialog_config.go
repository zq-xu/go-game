package dialog

import (
	"github.com/zq-xu/gotools/configx"

	"github.com/zq-xu/go-game/internal/dungeon/resources"
	"github.com/zq-xu/go-game/pkg/graphics/images/imagekit"
)

type DialogConfig struct {
	Icon string
	Box  string
}

var DialogCfg DialogConfig

func GetDialogConfig() *DialogConfig {
	return &DialogCfg
}

func init() {
	configx.RegisterByFile("dialog", &DialogCfg, configx.DefaultSetupFunc)
}

func NewDialogBox() (imagekit.Image, error) {
	cfg := GetDialogConfig()
	return resources.NewDungeonImage(cfg.Box)
}

func NewDialogIcon() (imagekit.Image, error) {
	cfg := GetDialogConfig()
	return resources.NewDungeonImage(cfg.Icon)
}
