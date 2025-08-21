package status

const (
	RunningStatus Status = iota
	TerminatingStatus
	NextStatus
	BackStatus
	SuccessStatus
	FailStatus
	ConfirmStatus
)

type Status int
