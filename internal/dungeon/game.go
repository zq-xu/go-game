package dungeon

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rotisserie/eris"

	"github.com/zq-xu/go-game/internal/dungeon/tiledmap"
)

const (
	mapImg  = "assets/tiled/map1.png"
	mapPath = "assets/tiled/map1.tmx"

	mapWidth  = 800
	mapHeight = 800
)

type game struct {
	tMap tiledmap.TiledMap
}

// StartGame
func StartGame() error {
	ebiten.SetWindowSize(mapWidth, mapHeight)
	ebiten.SetWindowTitle("Tiled map demo")

	g, err := NewGame()
	if err != nil {
		return eris.Wrap(err, "failed to initialize game")
	}

	return ebiten.RunGame(g)
}

func NewGame() (ebiten.Game, error) {
	var err error
	g := &game{}
	g.tMap, err = tiledmap.NewTiledMap(mapImg, mapPath)
	return g, err
}

func (g *game) Update() error { return nil }

func (g *game) Draw(screen *ebiten.Image) {
	g.tMap.DrawBackground(screen, 0, 0)
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
