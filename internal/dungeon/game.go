package dungeon

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rotisserie/eris"

	"github.com/zq-xu/go-game/internal/dungeon/actor"
	"github.com/zq-xu/go-game/internal/dungeon/config"
	"github.com/zq-xu/go-game/internal/dungeon/tiledmap"
	"github.com/zq-xu/go-game/pkg/event/collision"
)

type game struct {
	tMap  tiledmap.TiledMap
	actor actor.Actor

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

	g.tMap, err = tiledmap.NewTiledMap(config.MapImg, config.MapPath)
	if err != nil {
		return nil, eris.Wrap(err, "failed to load tiledmap")
	}

	g.actor, err = actor.NewActor()
	if err != nil {
		return nil, eris.Wrap(err, "failed to load actor")
	}

	g.collisionSpace = collision.NewCollisionSpace(
		g.tMap.Width(), g.tMap.Height(),
		g.tMap.TileWidth(), g.tMap.TileHeight())

	g.collisionSpace.AddObject(g.tMap.CollisionObjects()...)
	g.collisionSpace.AddObject(g.actor.CollisionObject())

	return g, err
}

func (g *game) Update() error {
	g.actor.Update()
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	g.tMap.DrawBackground(screen, 0, 0)
	g.actor.Draw(screen)
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
