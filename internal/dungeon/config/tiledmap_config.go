package config

import (
	"github.com/zq-xu/go-game/assets/dungeon"
)

const (
	MapImg  = "resources/map1.png"
	MapPath = "assets/dungeon/map1.tmx"

	MapWidth  = 800
	MapHeight = 800
)

type TiledMapConfig struct {
	BgImgPath, MapPath string
	Gifs               map[string]*GifConfig
}

type GifConfig struct {
	ImgPaths []string
	X, Y     float64
}

func GetTiledmapConfig() *TiledMapConfig {
	return &TiledMapConfig{
		BgImgPath: MapImg,
		MapPath:   MapPath,
		Gifs: map[string]*GifConfig{
			"wall1-torch1": {
				ImgPaths: []string{
					dungeon.GetDungeonImagePath("gif/torch/torch_1.png"),
					dungeon.GetDungeonImagePath("gif/torch/torch_2.png"),
					dungeon.GetDungeonImagePath("gif/torch/torch_3.png"),
					dungeon.GetDungeonImagePath("gif/torch/torch_4.png"),
				},
				X: 200,
				Y: 200,
			},
			"wall1-torch2": {
				ImgPaths: []string{
					dungeon.GetDungeonImagePath("gif/torch/torch_1.png"),
					dungeon.GetDungeonImagePath("gif/torch/torch_2.png"),
					dungeon.GetDungeonImagePath("gif/torch/torch_3.png"),
					dungeon.GetDungeonImagePath("gif/torch/torch_4.png"),
				},
				X: 230,
				Y: 200,
			},
			"wall2-candle2": {
				ImgPaths: []string{
					dungeon.GetDungeonImagePath("gif/candleA/candleA_01.png"),
					dungeon.GetDungeonImagePath("gif/candleA/candleA_02.png"),
					dungeon.GetDungeonImagePath("gif/candleA/candleA_03.png"),
					dungeon.GetDungeonImagePath("gif/candleA/candleA_04.png"),
				},
				X: 580,
				Y: 500,
			},
		},
	}
}
