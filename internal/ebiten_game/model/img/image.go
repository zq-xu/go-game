package img

import (
	"bytes"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/zq-xu/2d-game/internal/ebiten_game/config"
)

type Image struct {
	Image *ebiten.Image

	Width  int
	Height int

	X float64 // x position
	Y float64 // y position

	sc *config.ScreenConfig
}

func NewImage(imgByte []byte, sc *config.ScreenConfig) *Image {
	img, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(imgByte))
	if err != nil {
		log.Fatal(err)
	}

	return &Image{
		Image:  img,
		Width:  img.Bounds().Dx(),
		Height: img.Bounds().Dy(),
		sc:     sc,
	}
}

// func NewGopherImg() *Image {
// 	p := path.Join(config.Cfg.ImageRootPath, "gopher.png")
// 	return NewImage(p)
// }

func (i *Image) SetX(x float64) {
	if x < 0 {
		i.X = 0
		return
	}

	maxX := float64(i.sc.ScreenWidth - i.Width)
	if x > maxX {
		i.X = maxX
		return
	}

	i.X = x
}

func (i *Image) SetY(y float64) {
	if y < 0 {
		i.Y = 0
		return
	}

	maxY := float64(i.sc.ScreenHeight - i.Height)
	if y > maxY {
		i.Y = maxY
		return
	}

	i.Y = y
}

func (i *Image) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(i.X, i.Y)
	screen.DrawImage(i.Image, op)
}
