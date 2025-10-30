package collision

import (
	"github.com/solarlune/resolv"
	"github.com/zq-xu/gotools/logx"
)

type Space interface {
	AddObject(objs ...Object)

	Debug()
}

type resolvSpace struct {
	space *resolv.Space

	objRecord map[uint32]string
}

// NewResolvCollision
func NewResolvSpace(spaceWidth, spaceHeight, cellWidth, cellHeight int) Space {
	rc := &resolvSpace{objRecord: make(map[uint32]string, 0)}
	logx.Logger.Debug("space size:", spaceWidth, spaceHeight, cellWidth, cellHeight)
	rc.space = resolv.NewSpace(spaceWidth, spaceHeight, cellWidth, cellHeight)
	return rc
}

func (rc *resolvSpace) AddObject(objs ...Object) {
	for _, v := range objs {
		rc.objRecord[v.object().ID()] = v.Name()
		rc.space.Add(v.object())
		v.setSpace(rc)
	}
}

func (rc *resolvSpace) Debug() {
	rc.space.ForEachShape(func(shape resolv.IShape, index, maxCount int) bool {
		logx.Logger.Debugf("resolv space %d shape range: (x=%.1f, y=%.1f, width=%.1f, height=%.1f)\n",
			shape.ID(),
			shape.Position().X-shape.Bounds().Width()/2,
			shape.Position().Y-shape.Bounds().Height()/2,
			shape.Bounds().Width(),
			shape.Bounds().Height(),
		)
		return true
	})
}
