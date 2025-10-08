package dungeon

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rotisserie/eris"

	"github.com/zq-xu/go-game/internal/dungeon/characters"
	"github.com/zq-xu/go-game/internal/dungeon/config"
	"github.com/zq-xu/go-game/internal/dungeon/core/actor"
	"github.com/zq-xu/go-game/internal/dungeon/tiledmap"
	"github.com/zq-xu/go-game/pkg/event/collision"
)

type game struct {
	tMap tiledmap.TiledMap

	characters     []actor.Actor
	collisionSpace collision.Space
}

// StartGame
func StartGame() error {
	ebiten.SetWindowSize(config.MapWidth, config.MapHeight)
	ebiten.SetWindowTitle("Dungeon")

	g, err := NewGame()
	if err != nil {
		return eris.Wrap(err, "failed to initialize game")
	}

	return ebiten.RunGame(g)
}

func NewGame() (ebiten.Game, error) {
	var err error
	g := &game{}

	g.tMap, err = tiledmap.NewTiledMap(config.GetTiledmapConfig())
	if err != nil {
		return nil, eris.Wrap(err, "failed to load tiledmap")
	}

	g.characters, err = characters.NewCharacters()
	if err != nil {
		return nil, eris.Wrap(err, "failed to load characters")
	}

	g.collisionSpace = collision.NewResolvSpace(
		g.tMap.Width(), g.tMap.Height(),
		g.tMap.TileWidth(), g.tMap.TileHeight())

	g.collisionSpace.AddObject(g.tMap.CollisionObjects()...)
	for _, v := range g.characters {
		g.collisionSpace.AddObject(v)
	}

	return g, nil
}

func (g *game) Update() error {
	for _, v := range g.characters {
		v.Update()
	}
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	g.tMap.Draw(screen, 0, 0)

	for _, v := range g.characters {
		v.Draw(screen)
	}
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
