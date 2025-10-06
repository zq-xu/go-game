package resources

import (
	"github.com/zq-xu/go-game/assets/dungeon"
	"github.com/zq-xu/go-game/pkg/graphics/images/imagekit"
)

// NewDungeonImage
func NewDungeonImage(imgPath string) (imagekit.Image, error) {
	return imagekit.NewImageFromEmbed(&dungeon.EmbeddedDungeon, imgPath)
}

// NewDungeonIImageList
func NewDungeonIImageList(imgPaths []string) ([]imagekit.Image, error) {
	return imagekit.NewImageListFromEmbed(&dungeon.EmbeddedDungeon, imgPaths)
}
