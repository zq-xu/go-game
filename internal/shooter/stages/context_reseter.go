package stages

type StageReseter interface {
	Reset()
}

type defaultStageReseter struct{}

func newDefaultStageReseter() StageReseter { return &defaultStageReseter{} }

func (d *defaultStageReseter) Reset() {}
