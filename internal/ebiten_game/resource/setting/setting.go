package setting

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/2d-game/internal/ebiten_game/config"
)

type Setting struct {
	DebugMode bool

	ScreenWidth  int
	ScreenHeight int

	ShipXSpeedFactor  float64
	ShipYSpeedFactor  float64
	BulletSpeedFactor float64

	UFOYSpeedFactor float64
}

func NewSetting(cfg *config.Config) *Setting {
	s := &Setting{
		ScreenWidth:       int(config.BaseScreenWidth),
		ScreenHeight:      int(config.BaseScreenHeight),
		ShipXSpeedFactor:  config.BaseShipXSpeedFactor,
		ShipYSpeedFactor:  config.BaseShipYSpeedFactor,
		BulletSpeedFactor: config.BaseBulletSpeedFactor,
		UFOYSpeedFactor:   config.BaseUFOYSpeedFactor,
	}

	s.adaptForWindowsWidth(cfg.ScreenConfig.ScreenWidth)
	s.adaptForWindowsHeight(cfg.ScreenConfig.ScreenHeight)

	return s
}

func (s *Setting) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return s.ScreenWidth, s.ScreenHeight
}

func (s *Setting) Listen() {
	s.listenWindows()

}

func (s *Setting) listenWindows() {
	w, h := ebiten.WindowSize()
	s.adaptForWindowsWidth(w)
	s.adaptForWindowsHeight(h)
}

func (s *Setting) adaptForWindowsWidth(w int) {
	if w == s.ScreenWidth {
		return
	}

	s.ScreenWidth = w
	xTimes := float64(s.ScreenWidth) / config.BaseScreenWidth

	s.ShipXSpeedFactor = config.BaseShipXSpeedFactor * xTimes
}

func (s *Setting) adaptForWindowsHeight(h int) {
	if h == s.ScreenHeight {
		return
	}

	s.ScreenHeight = h
	yTimes := float64(s.ScreenHeight) / config.BaseScreenHeight

	s.ShipYSpeedFactor = config.BaseShipYSpeedFactor * yTimes
	s.BulletSpeedFactor = config.BaseBulletSpeedFactor * yTimes
}
