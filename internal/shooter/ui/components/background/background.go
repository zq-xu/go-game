package background

import (
	"github.com/ebitenui/ebitenui/image"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rotisserie/eris"

	"github.com/zq-xu/go-game/assets"
	"github.com/zq-xu/go-game/internal/shooter/ui/resources"
	"github.com/zq-xu/go-game/pkg/graphics"
	"github.com/zq-xu/go-game/pkg/graphics/images/background"
)

const backgroundColor = "131a22"

var (
	DeepStarrySkyBackgroundPath = assets.GetShooterImagePath("background/deep-starry-sky.jpg")
	MoonSurfaceBackgroundPath   = assets.GetShooterImagePath("background/moon-surface.jpg")
	NordwoodBackgroundPath      = assets.GetShooterImagePath("background/nordwood.jpg")
)

// NewDeepStarrySkyDownwardsBackground
func NewDeepStarrySkyDownwardsBackground() (ebiten.Game, error) {
	bi, err := resources.GetImage(DeepStarrySkyBackgroundPath)
	if err != nil {
		return nil, eris.Wrap(err, "get image failed.")
	}

	return background.NewYGradientDownwards(bi.Image(), bi.Width(), bi.Height()), nil
}

// NewMoonSurfaceDownwardsBackground
func NewMoonSurfaceDownwardsBackground() (ebiten.Game, error) {
	bi, err := resources.GetImage(MoonSurfaceBackgroundPath)
	if err != nil {
		return nil, eris.Wrap(err, "get image failed.")
	}
	return background.NewYGradientDownwards(bi.Image(), bi.Width(), bi.Height()), nil
}

// NewNordwoodDownwardsBackground
func NewNordwoodDownwardsBackground() (ebiten.Game, error) {
	bi, err := resources.GetImage(NordwoodBackgroundPath)
	if err != nil {
		return nil, eris.Wrap(err, "get image failed.")
	}
	return background.NewYGradientDownwards(bi.Image(), bi.Width(), bi.Height()), nil
}

// NewDefaultBackgroundImage
func NewDefaultBackgroundImage() *image.NineSlice {
	return image.NewNineSliceColor(graphics.HexToColor(backgroundColor))
}
