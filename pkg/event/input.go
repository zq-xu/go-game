package event

import (
	"fmt"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"

	"github.com/zq-xu/go-game/pkg/config"
	"github.com/zq-xu/go-game/pkg/metrics"
)

const InputName = "Input"

type Input struct {
	keys []ebiten.Key
}

func NewInput() *Input {
	return &Input{keys: make([]ebiten.Key, 0)}
}

func (i *Input) Update() {
	i.keys = inpututil.AppendPressedKeys(i.keys[:0])
}

func (i *Input) GetKeyString() string {
	var keyStrs []string
	var keyNames []string
	for _, k := range i.keys {
		keyStrs = append(keyStrs, k.String())
		if name := ebiten.KeyName(k); name != "" {
			keyNames = append(keyNames, name)
		}
	}

	return strings.Join(keyStrs, ", ") + "\n" + strings.Join(keyNames, ", ")
}

func (i *Input) Draw(screen *ebiten.Image, cfg *config.Config) {}

func (i *Input) DrawMetrics(screen *ebiten.Image, cfg *metrics.DrawConfig) {
	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(cfg.X), float64(cfg.Y))
	op.ColorScale.ScaleWithColor(cfg.Color)

	text.Draw(screen, fmt.Sprintf("KeyPress: %v", i.GetKeyString()), cfg.Face, op)
}
