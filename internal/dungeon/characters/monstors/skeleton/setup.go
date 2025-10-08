package skeleton

import (
	"fmt"
	"image"
	"math/rand"

	"github.com/rotisserie/eris"
	"github.com/zq-xu/go-game/internal/dungeon/config"
	"github.com/zq-xu/go-game/internal/dungeon/core/actor"
)

func NewSkeletonList() ([]actor.Actor, error) {
	list := make([]actor.Actor, 0)

	rects := generateRandomRects(5, 200, config.MapWidth, config.MapHeight)

	for k, v := range rects {
		s, err := NewSkeleton(&SkeletonConfig{
			Name:        fmt.Sprintf("Skeleton-%d", k),
			StartPoint:  &v.Min,
			StepLength:  1,
			ActiveRange: v,
		})
		if err != nil {
			return nil, eris.Wrap(err, "failed to new skeleton")
		}

		list = append(list, s)
	}
	return list, nil
}

// func generateRandomRect(n int, maxWidth, maxHeight int) *image.Rectangle {

// }

func generateRandomRects(n int, rectSize, maxWidth, maxHeight float64) []*image.Rectangle {
	rects := make([]*image.Rectangle, 0, n)

	for i := 0; i < n; i++ {
		x := rand.Intn(int(maxWidth - rectSize))
		y := rand.Intn(int(maxHeight - rectSize))

		rect := image.Rect(x, y, x+int(rectSize), y+int(rectSize))
		rects = append(rects, &rect)
	}

	return rects
}
