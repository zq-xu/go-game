package stages

type StageReseter interface {
	Reset() error
}

type defaultStageReseter struct{}

func newDefaultStageReseter() StageReseter { return &defaultStageReseter{} }

func (d *defaultStageReseter) Reset() error { return nil }
