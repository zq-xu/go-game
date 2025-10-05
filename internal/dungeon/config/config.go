package config

const (
	UpDirection Direction = iota
	DownDirection
	LeftDirection
	RightDirection
)

const DefaultDirection = DownDirection

const (
	IdleActorStatus ActorStatus = iota
	RunningActorStatus
	AttackActorStatus
)

type ActorStatus int
type Direction int

type Config struct{}
