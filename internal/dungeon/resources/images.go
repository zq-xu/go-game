package resources

import (
	"github.com/zq-xu/go-game/assets/dungeon"
	"github.com/zq-xu/go-game/pkg/graphics/images/imagekit"
)

// NewDungeonImage
func NewDungeonImage(path string) (imagekit.Image, error) {
	return imagekit.NewImageFromEmbed(&dungeon.EmbeddedDungeon, path)
}
