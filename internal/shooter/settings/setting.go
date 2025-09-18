package settings

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/zq-xu/go-game/pkg/config"
)

var globalSettings Settings

type Settings interface {
	DebugMode() bool
	ScreenWidth() int
	ScreenHeight() int

	ShipXSpeedFactor() float64
	ShipYSpeedFactor() float64
	BulletSpeedFactor() float64

	UFOYSpeedFactor() float64

	Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int)
	Listen()
}

type setting struct {
	debugMode bool

	screenWidth  int
	screenHeight int

	shipXSpeedFactor  float64
	shipYSpeedFactor  float64
	bulletSpeedFactor float64

	ufoYSpeedFactor float64
}

// GetSettings
func GetSettings() Settings {
	if globalSettings == nil {
		globalSettings = newSetting(config.Cfg)
	}
	return globalSettings
}

func newSetting(cfg *config.Config) Settings {
	s := &setting{
		screenWidth:       int(config.BaseScreenWidth),
		screenHeight:      int(config.BaseScreenHeight),
		shipXSpeedFactor:  BaseShipXSpeedFactor,
		shipYSpeedFactor:  BaseShipYSpeedFactor,
		bulletSpeedFactor: BaseBulletSpeedFactor,
		ufoYSpeedFactor:   BaseUFOYSpeedFactor,
	}

	s.adaptForWindowsWidth(cfg.ScreenConfig.ScreenWidth)
	s.adaptForWindowsHeight(cfg.ScreenConfig.ScreenHeight)

	return s
}

func (s *setting) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return s.screenWidth, s.screenHeight
}

func (s *setting) Listen() {
	s.listenWindows()
}

func (s *setting) listenWindows() {
	w, h := ebiten.WindowSize()
	s.adaptForWindowsWidth(w)
	s.adaptForWindowsHeight(h)
}

func (s *setting) adaptForWindowsWidth(w int) {
	if w == s.screenWidth {
		return
	}

	s.screenWidth = w
	xTimes := float64(s.screenWidth) / BaseScreenWidth

	s.shipXSpeedFactor = BaseShipXSpeedFactor * xTimes
}

func (s *setting) adaptForWindowsHeight(h int) {
	if h == s.screenHeight {
		return
	}

	s.screenHeight = h
	yTimes := float64(s.screenHeight) / BaseScreenHeight

	s.shipYSpeedFactor = BaseShipYSpeedFactor * yTimes
	s.bulletSpeedFactor = BaseBulletSpeedFactor * yTimes
}

func (s *setting) DebugMode() bool            { return s.debugMode }
func (s *setting) ScreenWidth() int           { return s.screenWidth }
func (s *setting) ScreenHeight() int          { return s.screenHeight }
func (s *setting) ShipXSpeedFactor() float64  { return s.shipXSpeedFactor }
func (s *setting) ShipYSpeedFactor() float64  { return s.shipYSpeedFactor }
func (s *setting) BulletSpeedFactor() float64 { return s.bulletSpeedFactor }
func (s *setting) UFOYSpeedFactor() float64   { return s.ufoYSpeedFactor }
