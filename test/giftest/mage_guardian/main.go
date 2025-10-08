package main

import (
	"github.com/zq-xu/go-game/pkg/utils"
	"github.com/zq-xu/go-game/test/giftest/base"
)

func main() {
	imgPaths, _ := utils.LoadFilesWithSuffix("./", ".png")
	base.ShowGifs(imgPaths...)
}
