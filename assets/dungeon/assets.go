package dungeon

import (
	"embed"
	"path"
)

//go:embed resources
var EmbeddedDungeon embed.FS

func GetDungeonImagePath(subPath string) string {
	return path.Join("resources", subPath)
}
