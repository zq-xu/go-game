package font

import (
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/sirupsen/logrus"
)

const (
	fontFaceRegular = "fonts/NotoSans-Regular.ttf"
	fontFaceBold    = "fonts/NotoSans-Bold.ttf"
)

var (
	notoSansTTF     []byte
	notoSansBoldTTF []byte
)

type notoSansFont struct {
	face     text.Face
	boldFace text.Face
}

func NewNotoSansFont(size float64) *notoSansFont {
	err := loadNotoSansFontTTF()
	if err != nil {
		logrus.Fatalf("load noto sans font ttf failed. %s", err)
		return nil
	}

	fl := &notoSansFont{}
	fl.face, _ = LoadFont(size, notoSansTTF)
	fl.boldFace, _ = LoadFont(size, notoSansBoldTTF)
	return fl
}

func loadNotoSansFontTTF() error {
	var err error

	if len(notoSansTTF) == 0 {
		notoSansTTF, err = NewFontFromFile(fontFaceRegular)
		if err != nil {
			return err
		}
	}

	if len(notoSansBoldTTF) == 0 {
		notoSansBoldTTF, err = NewFontFromFile(fontFaceBold)
		if err != nil {
			return err
		}
	}

	return nil
}

func (fl *notoSansFont) Face() *text.Face     { return &fl.face }
func (fl *notoSansFont) BoldFace() *text.Face { return &fl.boldFace }
