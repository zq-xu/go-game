package event

import (
	"image"

	"github.com/zq-xu/2d-game/internal/ebiten_game/model/img"
)

func CheckCollision(imgA, imgB *img.Image) bool {
	rectA := image.Rect(int(imgA.X), int(imgA.Y), int(imgA.X)+imgA.Width, int(imgA.Y)+imgA.Height)
	rectB := image.Rect(int(imgB.X), int(imgB.Y), int(imgB.X)+imgB.Width, int(imgB.Y)+imgB.Height)
	return rectA.Overlaps(rectB)
}
