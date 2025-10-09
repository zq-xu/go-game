package dialog

import (
	"github.com/zq-xu/go-game/internal/dungeon/config"
	"github.com/zq-xu/go-game/internal/dungeon/resources"
	"github.com/zq-xu/go-game/pkg/graphics/images/imagekit"
)

func NewDialogIcon() (imagekit.Image, error) {
	cfg := config.GetDialogConfig()
	return resources.NewDungeonImage(cfg.Icon)
}
