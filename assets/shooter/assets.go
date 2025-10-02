package shooter

import (
	"embed"
	"path"
)

func GetShooterImagePath(subPath string) string {
	return path.Join("images", subPath)
}

func GetShooterFontPath(subPath string) string {
	return path.Join("fonts", subPath)
}

//go:embed images
var EmbeddedShooterImage embed.FS

//go:embed fonts
var EmbeddedShooterFont embed.FS
