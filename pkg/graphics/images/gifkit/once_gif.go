package gifkit

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/pkg/graphics/images/imagekit"
)

type onceGif struct {
	drawing bool
	*gifBase
}

func NewOnceGif(imgs ...imagekit.Image) Gif {
	return &onceGif{
		gifBase: newGifBase(imgs...),
	}
}

// func NewOnceGifFromEmbed(embedFS *embed.FS, imgPaths []string) (Gif, error) {
// 	g, err := newGifBaseFromEmbed(embedFS, imgPaths)
// 	if err != nil {
// 		return nil, err
// 	}

// 	o := &onceGif{
// 		gifBase: g,
// 	}

// 	return o, nil
// }

func (o *onceGif) IsDrawing() bool { return o.drawing }

func (o *onceGif) MoveToStart() { o.gifBase.MoveToStart() }

func (o *onceGif) Draw(screen *ebiten.Image, x, y float64) {
	o.gifBase.draw(screen, x, y)
	o.next()
}

func (o *onceGif) Image() imagekit.Image {
	defer o.next()
	return o.image()
}

func (o *onceGif) next() {
	maxIndex := len(o.gifBase.images) - 1
	switch {
	case o.gifBase.index < maxIndex:
		o.drawing = true
		o.gifBase.next()
	default:
		o.gifBase.index = maxIndex
		o.drawing = false
	}
}
