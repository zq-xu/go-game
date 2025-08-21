package stages

var stageSettings = make(map[StageName]GameStage)

// Register
func Register(name StageName, stage GameStage) {
	stageSettings[name] = stage
}

// GetGameStage
func GetGameStage(stageName StageName) GameStage {
	v, ok := stageSettings[stageName]
	if ok {
		return v
	}

	return nil
}
