package game

const (
	RunningStageStatus StageStatus = iota
	SuccessStageStatus
	FailStageStatus
	ReturnBackStageStatus
)

type StageStatus int
