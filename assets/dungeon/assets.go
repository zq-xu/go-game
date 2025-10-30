package dungeon

import (
	"embed"
	"path"
	"strings"
)

//go:embed resources
var EmbeddedDungeon embed.FS

func GetDungeonImagePath(subPath string) string {
	if strings.HasPrefix(subPath, "resources") {
		return subPath
	}

	return path.Join("resources", subPath)
}
