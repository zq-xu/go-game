package actor

import (
	"image"

	"github.com/zq-xu/go-game/pkg/event/input"
)

type Option struct {
	Name string

	StartPoint    *image.Point
	Width, Height float64

	AttackImagePaths map[input.Direction]string
	MovingImagePaths map[input.Direction]string
	IdleImagePaths   map[input.Direction]string

	StepLength float64
}
