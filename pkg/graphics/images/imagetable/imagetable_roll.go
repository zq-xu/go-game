package imagetable

import "github.com/zq-xu/go-game/pkg/graphics/images/imagekit"

type RollImages interface {
	Next()
	Image() imagekit.Image
}

type rollImageTable struct {
	imageTable ImageTable

	length int
	index  int
}

// NewRollImages
func NewRollImages(img imagekit.Image) RollImages {
	it := NewDungeonImageTable(img)
	// it.LogBoxes()

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
