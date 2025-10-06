package collision

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"

	"github.com/zq-xu/go-game/pkg/logs"
)

type Object interface {
	// The point from topLeft
	LeftTop() (float64, float64)

	// The point from center
	Center() (float64, float64)

	Move(x, y float64)
	MoveByKey(key ebiten.Key, stepLength float64)

	Name() string

	object() *resolv.ConvexPolygon
	setSpace(s *resolvSpace)
}

type resolvObject struct {
	w, h float64

	obj *resolv.ConvexPolygon

	name string

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

func (o *resolvObject) LeftTop() (float64, float64) {
	c := o.obj.Position()
	return c.X - o.w/2, c.Y - o.h/2
}

func (o *resolvObject) Center() (float64, float64) {
	c := o.obj.Position()
	return c.X, c.Y
}

func (o *resolvObject) Y() float64 { return o.obj.Position().Y - o.h/2 }

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

func (o *resolvObject) Name() string { return o.name }

func (o *resolvObject) object() *resolv.ConvexPolygon { return o.obj }

func (o *resolvObject) setSpace(s *resolvSpace) { o.space = s }
