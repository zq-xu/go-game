package event

import (
	"image"

	"github.com/zq-xu/2d-game/pkg/graphics"
)

func CheckCollision(imgA, imgB *graphics.ImageEntity) bool {
	rectA := image.Rect(int(imgA.X), int(imgA.Y), int(imgA.X)+imgA.Img.Width, int(imgA.Y)+imgA.Img.Height)
	rectB := image.Rect(int(imgB.X), int(imgB.Y), int(imgB.X)+imgB.Img.Width, int(imgB.Y)+imgB.Img.Height)
	return rectA.Overlaps(rectB)
}
