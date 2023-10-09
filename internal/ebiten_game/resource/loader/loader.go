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
		FontLoader:  NewFontLoader(),
		ImageLoader: NewImageLoader(),
	}
}
