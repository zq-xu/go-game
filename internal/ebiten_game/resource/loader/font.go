package loader

import "github.com/zq-xu/2d-game/internal/ebiten_game/resource/loader/font"

const (
	DPI = 72

	BigTitleFaceSize = 36
	TitleFaceSize    = 24
	FaceSize         = 12
	SmallFaceSize    = 8
)

var (
	FontCfg = &font.Config{
		DPI:              DPI,
		BigTitleFaceSize: BigTitleFaceSize,
		TitleFaceSize:    TitleFaceSize,
		FaceSize:         FaceSize,
		SmallFaceSize:    SmallFaceSize,
	}
)

type FontLoader struct {
	font.Interface
}

func NewFontLoader() *FontLoader {
	fl := &FontLoader{
		Interface: font.NewNotoSansFontLoader(FontCfg),
	}

	return fl
}
