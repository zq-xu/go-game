package actors

import (
	"image"

	"github.com/zq-xu/go-game/internal/shooter/stages/gaming/gamerun/actors/entity"
)

func CheckCollision(imgA, imgB *entity.ImageEntity) bool {
	rectA := image.Rect(int(imgA.X), int(imgA.Y), int(imgA.X)+imgA.Img.Width(), int(imgA.Y)+imgA.Img.Height())
	rectB := image.Rect(int(imgB.X), int(imgB.Y), int(imgB.X)+imgB.Img.Width(), int(imgB.Y)+imgB.Img.Height())
	return rectA.Overlaps(rectB)
}
