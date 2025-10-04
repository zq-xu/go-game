package imagetable

import "github.com/zq-xu/go-game/pkg/graphics/images/imagekit"

type RollImages interface {
	Next()
	MoveToStart()
	CurrentIndex() int

	Image() imagekit.Image
	Range(fn func(index int, img imagekit.Image))
}

type rollImageTable struct {
	imageTable ImageTable

	length int
	index  int
}

// NewRollImages
func NewRollImages(name string, img imagekit.Image) RollImages {
	it := NewDungeonImageTable(name, img)
	it.LogBoxes()

	return &rollImageTable{
		imageTable: it,
		length:     it.Length(),
		index:      0,
	}
}

func (ri *rollImageTable) Next() {
	ri.index = (ri.index + 1) % ri.length

}

func (ri *rollImageTable) Image() imagekit.Image {
	return ri.imageTable.Images()[ri.index]
}

func (ri *rollImageTable) MoveToStart() {
	ri.index = 0
}

func (ri *rollImageTable) CurrentIndex() int { return ri.index }

func (ri *rollImageTable) Range(fn func(index int, img imagekit.Image)) {
	for k, v := range ri.imageTable.Images() {
		fn(k, v)
	}
}
