package font

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
	once sync.Once

	tt *opentype.Font

	TitleArcadeFont font.Face
	ArcadeFont      font.Face
	SmallArcadeFont font.Face
)

func init() {
	TitleArcadeFont, _ = NewFont(float64(config.Cfg.TitleFontSize))
	ArcadeFont, _ = NewFont(float64(config.Cfg.FontSize))
	SmallArcadeFont, _ = NewFont(float64(config.Cfg.SmallFontSize))
}

func initTT() {
	once.Do(func() {
		var err error

		tt, err = opentype.Parse(fonts.PressStart2P_ttf)
		if err != nil {
			log.Fatal(err)
		}
	})
}

func NewFont(size float64) (font.Face, error) {
	initTT()

	return opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(size),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}
