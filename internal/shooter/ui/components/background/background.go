package background

import (
	"github.com/ebitenui/ebitenui/image"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/assets"
	"github.com/zq-xu/go-game/pkg/graphics"
)

const backgroundColor = "131a22"

var (
	DeepStarrySkyBackgroundPath = assets.GetShooterImagePath("background/deep-starry-sky.jpg")
	MoonSurfaceBackgroundPath   = assets.GetShooterImagePath("background/moon-surface.jpg")
	NordwoodBackgroundPath      = assets.GetShooterImagePath("background/nordwood.jpg")
)

// NewDeepStarrySkyDownwardsBackground
func NewDeepStarrySkyDownwardsBackground() ebiten.Game {
	bi := graphics.GetImage(DeepStarrySkyBackgroundPath)
	return graphics.NewYGradientDownwards(bi.Image(), bi.Width(), bi.Height())
}

// NewMoonSurfaceDownwardsBackground
func NewMoonSurfaceDownwardsBackground() ebiten.Game {
	bi := graphics.GetImage(MoonSurfaceBackgroundPath)
	return graphics.NewYGradientDownwards(bi.Image(), bi.Width(), bi.Height())
}

// NewNordwoodDownwardsBackground
func NewNordwoodDownwardsBackground() ebiten.Game {
	bi := graphics.GetImage(NordwoodBackgroundPath)
	return graphics.NewYGradientDownwards(bi.Image(), bi.Width(), bi.Height())
}

// NewDefaultBackgroundImage
func NewDefaultBackgroundImage() *image.NineSlice {
	return image.NewNineSliceColor(graphics.HexToColor(backgroundColor))
}
