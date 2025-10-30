package config

const (
	IdleActorStatus ActorStatus = iota
	RunningActorStatus
	AttackActorStatus
)

type ActorStatus int
