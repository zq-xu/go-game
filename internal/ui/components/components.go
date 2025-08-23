package components

import (
	"github.com/zq-xu/go-game/internal/ui/components/background"
	"github.com/zq-xu/go-game/internal/ui/components/button"
	"github.com/zq-xu/go-game/internal/ui/components/text"
)

var (
	NewDefaultBackgroundImage           = background.NewDefaultBackgroundImage
	NewNordwoodDownwardsBackground      = background.NewNordwoodDownwardsBackground
	NewMoonSurfaceDownwardsBackground   = background.NewMoonSurfaceDownwardsBackground
	NewDeepStarrySkyDownwardsBackground = background.NewDeepStarrySkyDownwardsBackground

	NewCenterBoldText = text.NewCenterBoldText
	NewCenterText     = text.NewCenterText

	NewMenuButton    = button.NewMenuButton
	NewSettingButton = button.NewSettingButton
)
