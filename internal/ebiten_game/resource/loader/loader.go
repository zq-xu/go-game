package loader

import (
	"github.com/zq-xu/2d-game/internal/ebiten_game/config"
)

type Loader struct {
	Cfg *config.Config

	FontLoader *FontLoader

	ImageLoader *ImageLoader
}

func NewLoader() *Loader {
	return &Loader{
		Cfg:         config.Cfg,
		FontLoader:  NewFontLoader(),
		ImageLoader: NewImageLoader(),
	}
}
