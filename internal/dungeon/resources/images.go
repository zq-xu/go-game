package resources

import (
	"github.com/zq-xu/go-game/assets/dungeon"
	"github.com/zq-xu/go-game/pkg/graphics/images/imagekit"
)

// NewDungeonImage
func NewDungeonImage(imgPath string) (imagekit.Image, error) {
	return imagekit.NewImageFromEmbed(&dungeon.EmbeddedDungeon, dungeon.GetDungeonImagePath(imgPath))
}

// NewDungeonIImageList
func NewDungeonIImageList(imgPaths []string) ([]imagekit.Image, error) {
	list := make([]string, len(imgPaths))
	for k, v := range imgPaths {
		list[k] = dungeon.GetDungeonImagePath(v)
	}
	return imagekit.NewImageListFromEmbed(&dungeon.EmbeddedDungeon, list)
}
