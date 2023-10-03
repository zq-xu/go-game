package graphics

import (
	"bytes"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Image struct {
	Image *ebiten.Image

	Width  int
	Height int
}

func NewImage(imgByte []byte) (*Image, error) {
	img, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(imgByte))
	if err != nil {
		return nil, err
	}

	return &Image{
		Image:  img,
		Width:  img.Bounds().Dx(),
		Height: img.Bounds().Dy(),
	}, nil
}

type ImageEntity struct {
	Img *Image

	X float64 // x position
	Y float64 // y position

	ScreenWidth  int
	ScreenHeight int
}

func NewImageEntityWithImage(img *Image, w, h int) *ImageEntity {
	return &ImageEntity{
		Img:          img,
		ScreenWidth:  w,
		ScreenHeight: h,
	}
}

func NewImageEntity(imgByte []byte, w, h int) (*ImageEntity, error) {
	img, err := NewImage(imgByte)
	if err != nil {
		return nil, err
	}

	return NewImageEntityWithImage(img, w, h), nil
}

func (i *ImageEntity) SetX(x float64) {
	if x < 0 {
		i.X = 0
		return
	}

	maxX := float64(i.ScreenWidth - i.Img.Width)
	if x > maxX {
		i.X = maxX
		return
	}

	i.X = x
}

func (i *ImageEntity) SetY(y float64) {
	if y < 0 {
		i.Y = 0
		return
	}

	maxY := float64(i.ScreenHeight - i.Img.Height)
	if y > maxY {
		i.Y = maxY
		return
	}

	i.Y = y
}

func (i *ImageEntity) MoveLeft(dx float64) {
	i.X -= dx
}

func (i *ImageEntity) MoveRight(dx float64) {
	i.X += dx
}

func (i *ImageEntity) MoveUp(dy float64) {
	i.Y -= dy
}

func (i *ImageEntity) MoveDown(dy float64) {
	i.Y += dy
}

func (i *ImageEntity) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(i.X, i.Y)
	screen.DrawImage(i.Img.Image, op)
}

func (i *ImageEntity) IsUpOfScreen() bool {
	return i.Y <= -float64(i.Img.Height)
}

func (i *ImageEntity) IsDownOfScreen() bool {
	return i.Y >= float64(i.ScreenHeight)
}
