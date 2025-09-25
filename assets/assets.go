package assets

import (
	"embed"
	"path"
)

func GetShooterImagePath(subPath string) string {
	return path.Join("shooter/images", subPath)
}

func GetShooterFontPath(subPath string) string {
	return path.Join("shooter/fonts", subPath)
}

func GetDungeonImagePath(subPath string) string {
	return path.Join("dungeon/resources", subPath)
}

//go:embed dungeon/resources
var EmbeddedDungeon embed.FS

//go:embed shooter
var EmbeddedShooter embed.FS
