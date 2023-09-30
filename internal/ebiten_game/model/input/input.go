package input

import (
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"

	"github.com/zq-xu/2d-game/internal/ebiten_game/config"
	"github.com/zq-xu/2d-game/internal/ebiten_game/metric"
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

func (i *Input) Draw(screen *ebiten.Image, cfg *config.Config) {
	// text.Draw(screen, i.GetKeyString(), bitmapfont.Face, 4, 12, color.Black)
}

func (i *Input) DrawMetrics(screen *ebiten.Image, cfg *metric.DrawConfig) {
	text.Draw(screen, i.GetKeyString(), cfg.Face, cfg.X, cfg.Y, cfg.Color)
}
