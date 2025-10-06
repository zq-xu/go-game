package config

const (
	StepLength = 1
)

const (
	IdleActorStatus ActorStatus = iota
	RunningActorStatus
	AttackActorStatus
)

var (
	ActorWidth  float64 = 20
	ActorHeight float64 = 30
)

type ActorStatus int
