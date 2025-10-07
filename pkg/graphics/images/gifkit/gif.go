package gifkit

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/pkg/graphics/images/imagekit"
)

const (
	LeftTopDrawBeginning DrawBeginning = iota
	CenterDrawBeginning
	CenterBottomDrawBeginning
	RightBottomDrawBeginning
)

type DrawBeginning int
type Gif interface {
	Width() int
	Height() int

	Image() imagekit.Image

	Draw(screen *ebiten.Image, x, y float64)
	IsDrawing() bool
	MoveToStart()

	SetUpdateInterval(i int)
	SetDrawBeginning(b DrawBeginning)
}

type gifBase struct {
	updateInterval int
	counter        int

	index  int
	images []imagekit.Image

	drawBeginning DrawBeginning
}

func (g *gifBase) SetUpdateInterval(i int) { g.updateInterval = i }

func (g *gifBase) SetDrawBeginning(b DrawBeginning) { g.drawBeginning = b }

func (g *gifBase) MoveToStart() { g.index = 0 }

func newGifBase(imgs ...imagekit.Image) *gifBase {
	return &gifBase{
		index:  0,
		images: imgs,
	}
}

// func newGifBaseFromEmbed(embedFS *embed.FS, imgPaths []string) (*gifBase, error) {
// 	list, err := imagekit.NewImageListFromEmbed(embedFS, imgPaths)
// 	if err != nil {
// 		return nil, eris.Wrap(err, "new image list from embed failed")
// 	}

// 	return newGifBase(list...), nil
// }

func (g *gifBase) draw(screen *ebiten.Image, x, y float64) {
	img := g.image()

	switch g.drawBeginning {
	case CenterBottomDrawBeginning:
		x = x - float64(img.Width())/2
		y = y - float64(img.Height())
	case CenterDrawBeginning:
		x = x - float64(img.Width())/2
		y = y - float64(img.Height())/2
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	screen.DrawImage(img.Image(), op)
}

func (g *gifBase) image() imagekit.Image {
	return g.images[g.index]
}

func (g *gifBase) next() {
	if g.counter < g.updateInterval {
		g.counter++
		return
	}

	g.index = (g.index + 1) % len(g.images)
	g.counter = 0
}

func (g *gifBase) Width() int  { return g.images[0].Width() }
func (g *gifBase) Height() int { return g.images[0].Height() }
