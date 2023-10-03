package loader

import (
	"log"
	"sync"

	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/zq-xu/2d-game/internal/ebiten_game/config"
)

const dpi = 72

var (
	pressStart2POnce     sync.Once
	pressStart2PFontBase *opentype.Font
)

type FontLoader struct {
	TitleArcadeFont font.Face
	ArcadeFont      font.Face
	SmallArcadeFont font.Face
}

func NewFontLoader() *FontLoader {
	fl := &FontLoader{}
	fl.TitleArcadeFont, _ = NewPressStart2PFont(&opentype.FaceOptions{
		Size:    float64(config.Cfg.TitleFontSize),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	fl.ArcadeFont, _ = NewPressStart2PFont(
		&opentype.FaceOptions{
			Size:    float64(config.Cfg.FontSize),
			DPI:     dpi,
			Hinting: font.HintingFull,
		})

	fl.SmallArcadeFont, _ = NewPressStart2PFont(
		&opentype.FaceOptions{
			Size:    float64(config.Cfg.SmallFontSize),
			DPI:     dpi,
			Hinting: font.HintingFull,
		})

	return fl
}

func initPressStart2PFontBase() {
	pressStart2POnce.Do(func() {
		var err error

		pressStart2PFontBase, err = opentype.Parse(fonts.PressStart2P_ttf)
		if err != nil {
			log.Fatal(err)
		}
	})
}

func NewPressStart2PFont(ops *opentype.FaceOptions) (font.Face, error) {
	initPressStart2PFontBase()
	return opentype.NewFace(pressStart2PFontBase, ops)
}

// var (
// 	mPlus1pRegularOnce     sync.Once
// 	mPlus1pRegularFontBase *opentype.Font
// )

// func initMPlus1pRegularFontBase() {
// 	mPlus1pRegularOnce.Do(func() {
// 		var err error

// 		mPlus1pRegularFontBase, err = opentype.Parse(fonts.MPlus1pRegular_ttf)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	})
// }

// func NewMPlus1pRegularFont(ops *opentype.FaceOptions) (font.Face, error) {
// 	initMPlus1pRegularFontBase()
// 	return opentype.NewFace(mPlus1pRegularFontBase, ops)
// }
