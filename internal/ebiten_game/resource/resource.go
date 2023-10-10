package resource

import (
	"github.com/zq-xu/2d-game/internal/ebiten_game/config"
	"github.com/zq-xu/2d-game/internal/ebiten_game/resource/loader"
	"github.com/zq-xu/2d-game/internal/ebiten_game/resource/setting"
	"github.com/zq-xu/2d-game/internal/ebiten_game/resource/ui"
)

type Resource struct {
	Cfg *config.Config

	loader.Loader

	ui.UIResource

	setting.Setting
}

func NewResource() *Resource {
	rsc := &Resource{}

	rsc.Cfg = config.Cfg

	rsc.Loader = *loader.NewLoader()

	rsc.UIResource = *ui.NewUIResource(rsc.Cfg, &rsc.Loader)

	rsc.Setting = *setting.NewSetting(rsc.Cfg)

	return rsc
}
