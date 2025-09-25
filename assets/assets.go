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

//go:embed shooter/images
var EmbeddedImages embed.FS

//go:embed shooter/fonts
var EmbeddedFonts embed.FS
