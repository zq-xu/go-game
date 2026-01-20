package stage

type StageController interface {
	Stage

	AppendStage(s Stage)
	SetStage(name StageName)
	UseDefault()
}

type stageController struct {
	stageSettings map[StageName]Stage

	// The default game stage
	defaultStage Stage

	// The current game stage
	Stage
}

// NewStageController
func NewStageController(s Stage) StageController {
	sc := &stageController{
		stageSettings: map[StageName]Stage{s.StageName(): s},
		Stage:         s,
	}

	return sc
}

func (sc *stageController) AppendStage(s Stage) {
	sc.stageSettings[s.StageName()] = s
}

func (sc *stageController) SetStage(name StageName) {
	if sc.Stage.StageName() == name {
		return
	}

	sc.Stage = sc.stageSettings[name]
}

func (sc *stageController) UseDefault() {
	sc.Stage = sc.defaultStage
}
