package config

const (
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

var TiledMapCfg TiledMapConfig

func GetTiledmapConfig() *TiledMapConfig {
	return &TiledMapCfg
}
