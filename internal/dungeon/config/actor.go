package config

const (
	IdleActorStatus ActorStatus = iota
	RunningActorStatus
	AttackActorStatus
)

var (
	ActorWidth  float64 = 22
	ActorHeight float64 = 30
)

type ActorStatus int
