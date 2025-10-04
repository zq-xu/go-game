package collision

import (
	"github.com/solarlune/resolv"

	"github.com/zq-xu/go-game/pkg/logs"
)

type Object interface {
	// The point from topLeft
	X() float64
	Y() float64

	Move(x, y float64)
}

type resolvObject struct {
	w, h float64

	name string

	obj *resolv.ConvexPolygon

	space *resolvSpace
}

func NewResolvRectObject(name string, x, y, w, h float64) Object {
	return &resolvObject{
		name: name,
		w:    w,
		h:    h,

		obj: resolv.NewRectangleFromTopLeft(x, y, w, h),
	}
}

func (o *resolvObject) X() float64 { return o.obj.Position().X - o.w/2 }

func (o *resolvObject) Y() float64 { return o.obj.Position().Y - o.h/2 }

func (o *resolvObject) Move(x, y float64) {
	o.obj.Move(x, y)

	if o.isCollision() {
		o.obj.Move(-x, -y)
	}
}

func (o *resolvObject) isCollision() bool {
	return o.obj.IntersectionTest(resolv.IntersectionTestSettings{
		// Check only shapes that are near the rectangle (within 1 cell's margin)
		TestAgainst: o.obj.SelectTouchingCells(1).FilterShapes(),
		OnIntersect: func(set resolv.IntersectionSet) bool {
			logs.Logger.Debugf("Collision with %s", o.space.objRecord[set.OtherShape.ID()])
			return true
		},
	})
}
