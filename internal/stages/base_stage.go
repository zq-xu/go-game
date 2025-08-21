package stages

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/internal/settings"
)

type BaseStage struct {
	currentStage GameStage
	nextStage    GameStage

	shownAt time.Time
}

func NewBaseStage() *BaseStage {
	return &BaseStage{shownAt: time.Now()}
}

func (g *BaseStage) SetCurrentGameStage(s GameStage) { g.currentStage = s }
func (g *BaseStage) SetNexttGameStage(s GameStage)   { g.nextStage = s }

func (g *BaseStage) NextGameStage() GameStage {
	if g.nextStage == nil {
		return g.currentStage
	}

	tmp := g.nextStage
	g.nextStage = nil

	tmp.Reload()
	return tmp
}

func (g *BaseStage) Update() error { return nil }

func (g *BaseStage) Draw(screen *ebiten.Image) {}

func (g *BaseStage) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return settings.GetSettings().Layout(outsideWidth, outsideHeight)
}

func (g *BaseStage) Reload() { g.shownAt = time.Now() }

func (g *BaseStage) IsStable() bool {
	return time.Since(g.shownAt) > time.Millisecond*200
}
