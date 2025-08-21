package font

import (
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/gofont/gobold"
	"golang.org/x/image/font/gofont/goregular"
)

type goFont struct {
	face     text.Face
	boldFace text.Face
}

func NewGoFont(size float64) *goFont {
	face, _ := LoadFont(size, goregular.TTF)
	boldFace, _ := LoadFont(size, gobold.TTF)

	return &goFont{face: face, boldFace: boldFace}
}

func (gf *goFont) Face() *text.Face { return &gf.face }

func (gf *goFont) BoldFace() *text.Face { return &gf.boldFace }
