package data

const (
	WaitingStartStageResult StageResult = iota
	RunningStageResult
	SuccessStageResult
	FailStageResult
)

type StageResult int

type GameStage struct {
	Number int

	Result StageResult
}

func NewGameStage() *GameStage {
	return &GameStage{
		Number: 0,
		Result: WaitingStartStageResult,
	}
}
