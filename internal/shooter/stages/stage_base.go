package stages

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/internal/shooter/settings"
)

type BaseStage struct {
	ctx StageContext

	shownAt time.Time
}

func NewBaseStage(ctx StageContext) *BaseStage {
	return &BaseStage{
		ctx:     ctx,
		shownAt: time.Now(),
	}
}

func (g *BaseStage) Context() StageContext { return g.ctx }

func (g *BaseStage) Reset() { g.shownAt = time.Time{} }

func (g *BaseStage) Reload() { g.shownAt = time.Now() }

func (g *BaseStage) IsStable() bool {
	return time.Since(g.shownAt) > time.Millisecond*200
}

func (g *BaseStage) Update() error { return nil }

func (g *BaseStage) Draw(screen *ebiten.Image) {}

func (g *BaseStage) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return settings.GetSettings().Layout(outsideWidth, outsideHeight)
}
