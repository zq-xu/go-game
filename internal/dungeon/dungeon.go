package dungeon

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rotisserie/eris"

	"github.com/zq-xu/go-game/internal/dungeon/config"
	"github.com/zq-xu/go-game/internal/dungeon/entity"
	"github.com/zq-xu/go-game/internal/dungeon/entity/core/actor"
	"github.com/zq-xu/go-game/internal/dungeon/entity/core/tiledmap"
	"github.com/zq-xu/go-game/pkg/event/collision"
)

// Implement ebiten.Game
type game struct {
	tMap tiledmap.TiledMap

	entities       []actor.Actor
	collisionSpace collision.Space
}

// NewGame
func NewGame() (ebiten.Game, error) {
	var err error
	g := &game{}

	g.tMap, err = tiledmap.NewTiledMap(config.GetTiledmapConfig())
	if err != nil {
		return nil, eris.Wrap(err, "failed to load tiledmap")
	}

	g.entities, err = entity.NewEntities(g.tMap)
	if err != nil {
		return nil, eris.Wrap(err, "failed to load entities")
	}

	g.initialiseCollision()
	return g, nil
}

func (g *game) initialiseCollision() {
	g.collisionSpace = collision.NewResolvSpace(
		g.tMap.Width(), g.tMap.Height(),
		g.tMap.TileWidth(), g.tMap.TileHeight())

	g.collisionSpace.AddObject(g.tMap.CollisionObjects()...)
	for _, v := range g.entities {
		g.collisionSpace.AddObject(v)
	}
}

func (g *game) Update() error {
	for _, v := range g.entities {
		v.Update()
	}
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	g.tMap.Draw(screen, 0, 0)

	for _, v := range g.entities {
		v.Draw(screen)
	}
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
