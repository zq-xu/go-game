package mode

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"

	"github.com/zq-xu/2d-game/internal/ebiten_game/data"
	"github.com/zq-xu/2d-game/internal/ebiten_game/model/font"
)

var (
	SuccessText = "SUCCESS"
	FailureText = "FAILED"

	resultTextSet = map[data.StageResult]string{
		data.SuccessStageResult: SuccessText,
		data.FailStageResult:    FailureText,
	}
)

type GameOverModeHandler struct {
	gameData *data.GameData
	ls       ModeListener
}

func init() {
	Register(GameOverMode, NewGameOverModeHandler)
}

func NewGameOverModeHandler(gameData *data.GameData, ls ModeListener) GameModeHandler {
	return &GameOverModeHandler{
		gameData: gameData,
		ls:       ls,
	}
}

func (g *GameOverModeHandler) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		g.ls.SetMode(GamingMode)
	}

	return nil
}

func (g *GameOverModeHandler) Draw(screen *ebiten.Image) {
	titleTexts := []string{"GAME OVER", resultTextSet[g.gameData.GameStage.Result]}

	tittleHight := len(titleTexts) * 2 * g.gameData.Cfg.TitleFontSize
	textHeight := len(StartTexts) * 2 * g.gameData.Cfg.FontSize

	startY := (g.gameData.Cfg.ScreenHeight - tittleHight - textHeight) / 2

	for _, l := range titleTexts {
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

func (g *GameOverModeHandler) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.gameData.Layout(outsideWidth, outsideHeight)
}
