package collision

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"

	"github.com/zq-xu/go-game/pkg/event/input"
)

func (o *resolvObject) MoveByKey(key ebiten.Key, stepLength float64) {
	switch key {
	case ebiten.KeyUp:
		o.Move(0, -stepLength)
	case ebiten.KeyDown:
		o.Move(0, stepLength)
	case ebiten.KeyLeft:
		o.Move(-stepLength, 0)
	case ebiten.KeyRight:
		o.Move(stepLength, 0)
	}
}

func (o *resolvObject) Move(x, y float64) {
	o.obj.Move(x, y)

	mtvList, collided := o.getMtvForCollision()
	if !collided {
		return
	}

	for _, v := range mtvList {
		o.obj.Move(v.X, v.Y)
	}
}

func (o *resolvObject) getMtvForCollision() ([]mtv, bool) {
	mtvList := make([]mtv, 0)

	for _, shape := range o.obj.SelectTouchingCells(1).FilterShapes().Shapes() {
		result := o.obj.Intersection(shape)
		if result.IsEmpty() {
			continue
		}

		mtv := computeMTV(o.obj, shape)
		mtvList = append(mtvList, mtv)

		// logs.Logger.Debugf("Collision with %s, direction: %+v,mtv: %+v",
		// 	o.space.objRecord[shape.ID()], getDirectionFromIntersectionSet(&result), mtv)
		return mtvList, true
	}

	return nil, false
}

func getDirectionFromIntersectionSet(set *resolv.IntersectionSet) []input.Direction {
	list := make([]input.Direction, 0)
	for _, v := range set.Intersections {
		nx := v.Normal.X
		ny := v.Normal.Y

		// for other shape, it's collision on right side,
		// for the object, it's on the left side
		if nx > 0 {
			list = append(list, input.RightDirection)
		}
		if nx < 0 {
			list = append(list, input.LeftDirection)
		}
		if ny > 0 {
			list = append(list, input.DownDirection)
		}
		if ny < 0 {
			list = append(list, input.UpDirection)
		}
	}
	return list
}
