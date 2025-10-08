package collision

import (
	"github.com/solarlune/resolv"

	"github.com/zq-xu/go-game/pkg/event/input"
)

type Object interface {
	// The point from topLeft
	LeftTop() (float64, float64)

	// The point from center
	Center() (float64, float64)

	Move(x, y float64)
	MoveDirection(d input.Direction)

	Name() string

	object() *resolv.ConvexPolygon
	setSpace(s *resolvSpace)
}

type resolvObject struct {
	w, h float64

	obj *resolv.ConvexPolygon

	name string

	space *resolvSpace

	stepLength float64
}

type objectOption func(o *resolvObject)

func WithStepLength(si float64) objectOption {
	return func(o *resolvObject) {
		o.stepLength = si
	}
}

func NewResolvRectObject(name string, x, y, w, h float64, opts ...objectOption) Object {
	o := &resolvObject{
		name:       name,
		w:          w,
		h:          h,
		stepLength: 1,

		obj: resolv.NewRectangleFromTopLeft(x, y, w, h),
	}

	for _, opt := range opts {
		opt(o)
	}

	return o
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
