package metrics

import (
	"fmt"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type PoolList interface {
	Add(name string, p MetricsPool)
	Draw(screen *ebiten.Image)
}

type poolList struct {
	lock sync.RWMutex
	ps   []*poolItem
}

type poolItem struct {
	index int

	name string
	p    MetricsPool
}

func NewMetricsPoolList() *poolList {
	return &poolList{
		ps: make([]*poolItem, 0),
	}
}

func (l *poolList) Add(name string, p MetricsPool) {
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

func (l *poolList) Draw(screen *ebiten.Image) {
	startY := MetricLineHeight

	for k := range l.ps {
		l.ps[k].drawItemTitle(screen, startY)
		startY += MetricLineHeight

		l.ps[k].p.DrawMetrics(screen, poolItemStartX, startY)
		startY += l.ps[k].p.Length()*MetricLineHeight*2 + MetricLineHeight*2
	}
}

func newPoolItem(index int, name string, p MetricsPool) *poolItem {
	return &poolItem{index, name, p}
}

func (pi *poolItem) drawItemTitle(screen *ebiten.Image, startY int) {
	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(poolItemTitleStartX), float64(startY))
	op.ColorScale.ScaleWithColor(defaultPoolItemDrawConfig.Color)

	text.Draw(screen, fmt.Sprintf("%s: ", pi.name), defaultPoolItemDrawConfig.Face, op)
}
