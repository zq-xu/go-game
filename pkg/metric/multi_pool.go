package metric

import (
	"fmt"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	poolItemTitleStartX = defaultStartX

	poolItemStartX = defaultStartX + 20
)

var (
	MultiPool = NewMetricPoolList()

	defaultPoolItemDrawConfig *DrawConfig
)

type PoolList struct {
	lock sync.RWMutex
	ps   []*poolItem
}

type poolItem struct {
	index int

	name string
	p    *Pool
}

func init() {
	defaultPoolItemDrawConfig = DefaultDrawConfig.Copy()
}

func newPoolItem(index int, name string, p *Pool) *poolItem {
	return &poolItem{index, name, p}
}

func NewMetricPoolList() *PoolList {
	return &PoolList{
		ps: make([]*poolItem, 0),
	}
}

func (l *PoolList) Add(name string, p *Pool) {
	l.lock.Lock()
	defer l.lock.Unlock()

	if p == nil {
		return
	}

	for k := range l.ps {
		if l.ps[k].name == name {
			l.ps[k].p = p
			return
		}
	}

	l.ps = append(l.ps, newPoolItem(len(l.ps), name, p))
}

func (l *PoolList) Draw(screen *ebiten.Image) {
	startY := MetricLineHeight

	for k := range l.ps {
		l.ps[k].drawItemTitle(screen, startY)
		startY += MetricLineHeight

		l.ps[k].p.drawMetrics(screen, poolItemStartX, startY)
		startY += len(l.ps[k].p.metricSet)*MetricLineHeight + 2*MetricLineHeight
	}
}

func (pi *poolItem) drawItemTitle(screen *ebiten.Image, startY int) {
	text.Draw(screen, fmt.Sprintf("%s: ", pi.name),
		defaultPoolItemDrawConfig.Face,
		poolItemTitleStartX,
		startY,
		defaultPoolItemDrawConfig.Color)
}
