package skeleton

import (
	"fmt"
	"image"
	"math/rand"

	"github.com/rotisserie/eris"
	"github.com/zq-xu/gotools/logx"

	"github.com/zq-xu/go-game/internal/dungeon/core/actor"
	"github.com/zq-xu/go-game/internal/dungeon/core/tiledmap"
)

func NewSkeletonList(tm tiledmap.TiledMap) ([]actor.Actor, error) {
	list := make([]actor.Actor, 0)
	rects := generateRandomRects(3, 200, float64(tm.Width()), float64(tm.Height()))

	logx.Logger.Info("Loading Skeletons")
	for k, v := range rects {
		s, err := NewSkeleton(&SkeletonConfig{
			Name:        fmt.Sprintf("Skeleton-%d", k),
			StartPoint:  &v.Min,
			ActiveRange: v,
		})
		if err != nil {
			return nil, eris.Wrap(err, "failed to new skeleton")
		}

		list = append(list, s)
	}
	logx.Logger.Info("Loaded Skeletons")

	return list, nil
}

func generateRandomRects(n int, rectSize, maxWidth, maxHeight float64) []*image.Rectangle {
	rects := make([]*image.Rectangle, 0, n)

	for range n {
		x := rand.Intn(int(maxWidth - rectSize))
		y := rand.Intn(int(maxHeight - rectSize))

		rect := image.Rect(x, y, x+int(rectSize), y+int(rectSize))
		rects = append(rects, &rect)
	}

	return rects
}
