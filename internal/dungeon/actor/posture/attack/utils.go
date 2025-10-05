package attack

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/internal/dungeon/config"
	"github.com/zq-xu/go-game/pkg/graphics/images/imagekit"
)

func getDrawBeginningOnRightBottom(img imagekit.Image, x, y float64) (float64, float64) {
	x = x + config.ActorWidth - float64(img.Width())
	y = y + config.ActorHeight - float64(img.Height())

	return x, y
}

func drawImage(screen *ebiten.Image, img imagekit.Image, x, y float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	screen.DrawImage(img.Image(), op)
}
