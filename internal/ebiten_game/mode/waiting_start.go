package mode

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"

	"github.com/zq-xu/2d-game/internal/ebiten_game/data"
	"github.com/zq-xu/2d-game/internal/ebiten_game/model/font"
)

var (
	TitleTexts = []string{"ALIEN INVASION", "AUTHOR: ZHIQIANG XU"}
	StartTexts = []string{"", "", "", "", "PRESS ENTER KEY TO START"}
)

type WaitingStartModeHandler struct {
	gameData *data.GameData
	ls       ModeListener
}

func init() {
	Register(WaitingStartMode, NewWaitingStartModeHandler)
}

func NewWaitingStartModeHandler(gameData *data.GameData, ls ModeListener) GameModeHandler {
	return &WaitingStartModeHandler{
		gameData: gameData,
		ls:       ls,
	}
}

func (g *WaitingStartModeHandler) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		g.ls.SetMode(GamingMode)
	}

	return nil
}

func (g *WaitingStartModeHandler) Draw(screen *ebiten.Image) {
	tittleHight := len(TitleTexts) * 2 * g.gameData.Cfg.TitleFontSize
	textHeight := len(StartTexts) * 2 * g.gameData.Cfg.FontSize

	startY := (g.gameData.Cfg.ScreenHeight - tittleHight - textHeight) / 2

	for _, l := range TitleTexts {
		x := (g.gameData.Cfg.ScreenWidth - len(l)*g.gameData.Cfg.TitleFontSize) / 2
		startY += 2 * g.gameData.Cfg.TitleFontSize
		text.Draw(screen, l, font.TitleArcadeFont, x, startY, color.White)
	}

	for _, l := range StartTexts {
		x := (g.gameData.Cfg.ScreenWidth - len(l)*g.gameData.Cfg.FontSize) / 2
		startY += 2 * g.gameData.Cfg.FontSize
		text.Draw(screen, l, font.ArcadeFont, x, startY, color.White)
	}
}

func (g *WaitingStartModeHandler) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.gameData.Layout(outsideWidth, outsideHeight)
}
