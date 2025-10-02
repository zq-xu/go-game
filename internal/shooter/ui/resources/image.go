package resources

import (
	"github.com/ebitenui/ebitenui/image"

	"github.com/zq-xu/go-game/assets"
	"github.com/zq-xu/go-game/pkg/graphics/images/imagekit"
	"github.com/zq-xu/go-game/pkg/graphics/images/imageloader"
)

var globalImageLoader imageloader.ImageLoader
var globalNineSliceImgLoader imageloader.NineSliceImgLoader

func init() {
	globalImageLoader = imageloader.NewimgLoader(&assets.EmbeddedShooter)
	globalNineSliceImgLoader = imageloader.NewNineSliceImgLoader(&assets.EmbeddedShooter)
}

// GetImage
func GetImage(imgPath string) (imagekit.Image, error) {
	return globalImageLoader.GetImage(imgPath)
}

// GetNineSliceSimpleImage
func GetNineSliceSimpleImage(path string, borderWidthHeight, centerWidthHeight int) (*image.NineSlice, error) {
	return globalNineSliceImgLoader.GetNineSliceSimpleImage(path, borderWidthHeight, centerWidthHeight)
}

// GetNineSliceImage
func GetFixedNineSlice(path string) (*image.NineSlice, error) {
	return globalNineSliceImgLoader.GetFixedNineSlice(path)
}

// GetNineSliceImage
func GetNineSliceImage(path string, centerWidth int, centerHeight int) (*image.NineSlice, error) {
	return globalNineSliceImgLoader.GetNineSliceImage(path, centerWidth, centerHeight)
}
