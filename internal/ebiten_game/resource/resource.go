package resource

import (
	"github.com/zq-xu/2d-game/internal/ebiten_game/resource/loader"
	"github.com/zq-xu/2d-game/internal/ebiten_game/resource/ui"
)

type Resource struct {
	loader.Loader

	ui.UIResource
}

func NewResource() *Resource {
	rsc := &Resource{}

	rsc.Loader = *loader.NewLoader()
	rsc.UIResource = *ui.NewUIResource(&rsc.Loader)

	return rsc
}
