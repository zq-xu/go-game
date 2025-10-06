package input

import "github.com/hajimehoshi/ebiten/v2"

const (
	UpDirection Direction = iota + 1
	DownDirection
	LeftDirection
	RightDirection
)

const DefaultDirection = DownDirection
const unknownDirection = 0

var (
	directionSet = map[Direction]string{
		unknownDirection: "unknown",
		UpDirection:      "Up",
		DownDirection:    "Down",
		LeftDirection:    "Left",
		RightDirection:   "Right",
	}

	keyDirectionSet = map[ebiten.Key]Direction{
		ebiten.KeyDown:  DownDirection,
		ebiten.KeyUp:    UpDirection,
		ebiten.KeyLeft:  LeftDirection,
		ebiten.KeyRight: RightDirection,
	}
)

type Direction int

func (d Direction) String() string {
	s, ok := directionSet[d]
	if ok {
		return s
	}

	return directionSet[unknownDirection]
}

func (d Direction) Unknown() bool { return d == unknownDirection }

func GetDirection(key ebiten.Key) Direction { return keyDirectionSet[key] }

func RangeKeyDirections(fn func(key ebiten.Key, direction Direction)) {
	for k, v := range keyDirectionSet {
		fn(k, v)
	}
}
