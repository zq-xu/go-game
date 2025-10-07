package actor

import "github.com/zq-xu/go-game/pkg/event/input"

type Option struct {
	Name string

	StartX, StartY float64
	Width, Height  float64

	AttackImagePaths map[input.Direction]string
	MovingImagePaths map[input.Direction]string
	IdleImagePaths   map[input.Direction]string

	StepLength int
}
