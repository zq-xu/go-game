package font

import (
	"log"
	"sync"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"

	"github.com/zq-xu/go-game/pkg/graphics"
)

const (
	fontFaceRegular = "fonts/NotoSans-Regular.ttf"
	fontFaceBold    = "fonts/NotoSans-Bold.ttf"
)

type NotoSansFontLoader struct {
	once sync.Once

	regularTTF *truetype.Font
	boldTTF    *truetype.Font

	titleFace    font.Face
	face         font.Face
	bigTitleFace font.Face
}

func NewNotoSansFontLoader(cfg *Config) *NotoSansFontLoader {
	fl := &NotoSansFontLoader{}

	fl.titleFace = fl.newBoldFont(&truetype.Options{
		Size:    float64(cfg.TitleFaceSize),
		DPI:     cfg.DPI,
		Hinting: font.HintingFull,
	})

	fl.face = fl.newRegularFont(
		&truetype.Options{
			Size:    float64(cfg.FaceSize),
			DPI:     cfg.DPI,
			Hinting: font.HintingFull,
		})

	fl.bigTitleFace = fl.newBoldFont(
		&truetype.Options{
			Size:    float64(cfg.BigTitleFaceSize),
			DPI:     cfg.DPI,
			Hinting: font.HintingFull,
		})

	return fl
}

func (fl *NotoSansFontLoader) initNotoSansFontTTF() {
	fl.once.Do(func() {
		var err error

		regularData, err := graphics.NewFontFromFile(fontFaceRegular)
		if err != nil {
			log.Fatal(err)
		}

		fl.regularTTF, err = truetype.Parse(regularData)
		if err != nil {
			log.Fatal(err)
		}

		boldData, err := graphics.NewFontFromFile(fontFaceBold)
		if err != nil {
			log.Fatal(err)
		}

		fl.boldTTF, err = truetype.Parse(boldData)
		if err != nil {
			log.Fatal(err)
		}
	})
}

func (fl *NotoSansFontLoader) newRegularFont(ops *truetype.Options) font.Face {
	fl.initNotoSansFontTTF()
	return truetype.NewFace(fl.regularTTF, ops)
}

func (fl *NotoSansFontLoader) newBoldFont(ops *truetype.Options) font.Face {
	fl.initNotoSansFontTTF()
	return truetype.NewFace(fl.boldTTF, ops)
}

func (fl *NotoSansFontLoader) Face() font.Face {
	return fl.face
}

func (fl *NotoSansFontLoader) TitleFace() font.Face {
	return fl.titleFace
}

func (fl *NotoSansFontLoader) BigTitleFace() font.Face {
	return fl.bigTitleFace
}
