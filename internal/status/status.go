package status

const (
	RunningStatus Status = iota
	SuccessStatus
	FailStatus
)

type Status int
