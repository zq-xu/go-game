package metric

import "github.com/hajimehoshi/ebiten/v2"

var MultiPool = NewMetricPoolList()

type PoolList struct {
	ps []*Pool
}

func NewMetricPoolList() *PoolList {
	return &PoolList{
		ps: make([]*Pool, 0),
	}
}

func (l *PoolList) Add(p *Pool) {
	if p == nil {
		return
	}

	l.ps = append(l.ps, p)
}

func (l *PoolList) Draw(screen *ebiten.Image) {
	startY := metricLineHeight

	for k := range l.ps {
		l.ps[k].drawMetrics(screen, startY)
		startY += len(l.ps[k].metricSet)*metricLineHeight + 2*metricLineHeight
	}
}
