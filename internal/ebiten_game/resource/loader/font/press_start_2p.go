package font

import (
	"log"
	"sync"

	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type PressStart2PFontLoader struct {
	once sync.Once
	ttf  *opentype.Font

	titleFace    font.Face
	face         font.Face
	bigTitleFace font.Face
}

func NewPressStart2PFontLoader(cfg *Config) *PressStart2PFontLoader {
	fl := &PressStart2PFontLoader{}

	fl.titleFace, _ = fl.newFont(&opentype.FaceOptions{
		Size:    float64(cfg.TitleFaceSize),
		DPI:     cfg.DPI,
		Hinting: font.HintingFull,
	})

	fl.face, _ = fl.newFont(
		&opentype.FaceOptions{
			Size:    float64(cfg.FaceSize),
			DPI:     cfg.DPI,
			Hinting: font.HintingFull,
		})

	fl.bigTitleFace, _ = fl.newFont(
		&opentype.FaceOptions{
			Size:    float64(cfg.BigTitleFaceSize),
			DPI:     cfg.DPI,
			Hinting: font.HintingFull,
		})

	return fl
}

func (fl *PressStart2PFontLoader) initTTF() {
	fl.once.Do(func() {
		var err error

		fl.ttf, err = opentype.Parse(fonts.PressStart2P_ttf)
		if err != nil {
			log.Fatal(err)
		}
	})
}

func (fl *PressStart2PFontLoader) newFont(ops *opentype.FaceOptions) (font.Face, error) {
	fl.initTTF()
	return opentype.NewFace(fl.ttf, ops)
}

func (fl *PressStart2PFontLoader) Face() font.Face {
	return fl.face
}

func (fl *PressStart2PFontLoader) TitleFace() font.Face {
	return fl.titleFace
}

func (fl *PressStart2PFontLoader) BigTitleFace() font.Face {
	return fl.bigTitleFace
}
