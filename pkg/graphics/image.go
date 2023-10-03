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

	ScreenWidth  int
	ScreenHeight int

	X float64 // x position
	Y float64 // y position

	MinX float64
	MaxX float64

	MinY float64
	MaxY float64
}

func NewImageEntityWithImage(img *Image, w, h int) *ImageEntity {
	return &ImageEntity{
		Img:          img,
		ScreenWidth:  w,
		ScreenHeight: h,

		MaxX: float64(w - img.Width),
		MaxY: float64(h - img.Height),
	}
}

func NewImageEntity(imgByte []byte, w, h int) (*ImageEntity, error) {
	img, err := NewImage(imgByte)
	if err != nil {
		return nil, err
	}

	return NewImageEntityWithImage(img, w, h), nil
}

func (i *ImageEntity) SetXLimit(min, max float64) {
	i.MinX, i.MaxX = min, max
}

func (i *ImageEntity) SetYLimit(min, max float64) {
	i.MinY, i.MaxY = min, max
}

func (i *ImageEntity) UnlimitTop() {
	i.SetYLimit(-float64(i.Img.Height), i.MaxY)
}

func (i *ImageEntity) UnlimitBottom() {
	i.SetYLimit(i.MinY, float64(i.ScreenHeight))
}

func (i *ImageEntity) UnlimitLeft() {
	i.SetXLimit(-float64(i.Img.Width), i.MaxX)
}

func (i *ImageEntity) UnlimitRight() {
	i.SetXLimit(i.MinX, float64(i.ScreenWidth))
}

func (i *ImageEntity) SetX(x float64) {
	if x < i.MinX {
		i.X = i.MinX
		return
	}

	if x > i.MaxX {
		i.X = i.MaxX
		return
	}

	i.X = x
}

func (i *ImageEntity) SetY(y float64) {
	if y < i.MinY {
		i.Y = i.MinY
		return
	}

	if y > i.MaxY {
		i.Y = i.MaxY
		return
	}

	i.Y = y
}

func (i *ImageEntity) MoveLeft(dx float64) {
	i.SetX(i.X - dx)
}

func (i *ImageEntity) MoveRight(dx float64) {
	i.SetX(i.X + dx)
}

func (i *ImageEntity) MoveUp(dy float64) {
	i.SetY(i.Y - dy)
}

func (i *ImageEntity) MoveDown(dy float64) {
	i.SetY(i.Y + dy)
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
