package stages

import "github.com/zq-xu/go-game/internal/status"

type StageContext interface {
	// Reset for restarting the game
	Reset()

	// CurrentGameStage is used to switch the game stage
	SetCurrentGameStage(name StageName)
	CurrentGameStage() StageName

	// Status is used to set the game status, success, failure, etc.
	SetStatus(s status.Status)
	Status() status.Status

	// StageReseter is used to reset the game stage
	SetStageReseter(s StageReseter)
	StageReseter() StageReseter

	// TempDrawer is mainly used for pause stage as background drawer
	SetTempDrawer(s TempDrawer)
	TempDrawer() TempDrawer
}

type stageContext struct {
	currentStage StageName
	status       status.Status
	stageReseter StageReseter
	tempDrawer   TempDrawer
}

func NewStageContext() StageContext {
	s := &stageContext{}
	s.stageReseter = newDefaultStageReseter()

	s.Reset()
	return s
}

func (ctx *stageContext) Reset() {
	ctx.stageReseter.Reset()

	ctx.tempDrawer = newDefaultTempDrawer()

	ctx.status = status.RunningStatus
	ctx.currentStage = MenuStage
}

func (ctx *stageContext) SetCurrentGameStage(name StageName) { ctx.currentStage = name }

func (ctx *stageContext) CurrentGameStage() StageName { return ctx.currentStage }

func (ctx *stageContext) SetStatus(s status.Status) { ctx.status = s }

func (ctx *stageContext) Status() status.Status { return ctx.status }

func (ctx *stageContext) SetStageReseter(s StageReseter) { ctx.stageReseter = s }

func (ctx *stageContext) StageReseter() StageReseter { return ctx.stageReseter }

func (ctx *stageContext) SetTempDrawer(s TempDrawer) { ctx.tempDrawer = s }

func (ctx *stageContext) TempDrawer() TempDrawer { return ctx.tempDrawer }
