package input

import (
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	VailidKeys []ebiten.Key

	AttackKey = ebiten.KeySpace
)

func init() {
	VailidKeys = append(VailidKeys, AttackKey)

	for k := range keyDirectionSet {
		VailidKeys = append(VailidKeys, k)
	}
}
