package gaming

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/internal/stages/gaming/navbar"
)

var (
	NewNavbar = navbar.NewNavbar

	WithNavbarSettingButonOpts = navbar.WithNavbarSettingButonOpts
)

type Navbar interface {
	Update() error
	Draw(screen *ebiten.Image)
}
