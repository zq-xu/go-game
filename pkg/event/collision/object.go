package collision

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
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

func (o *resolvObject) Name() string { return o.name }

func (o *resolvObject) object() *resolv.ConvexPolygon { return o.obj }

func (o *resolvObject) setSpace(s *resolvSpace) { o.space = s }
