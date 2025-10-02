package imagetable

import (
	"fmt"
	"image"

	"github.com/zq-xu/go-game/pkg/graphics/images/imagekit"
)

type ImageTable interface {
	Images() []imagekit.Image
	Length() int

	// for debug
	LogBoxes()
}

type imageTable struct {
	boxes  []BoundingBox
	images []imagekit.Image
}

// NewDungeonImageTable
func NewDungeonImageTable(img imagekit.Image) ImageTable {
	it := &imageTable{boxes: NewBoundingBoxes(img.GoImage())}

	for _, v := range it.boxes {
		goImage := img.Image().SubImage(image.Rect(v.Xmin(), v.Ymin(), v.Xmax(), v.Ymax()))
		it.images = append(it.images, imagekit.NewBasicImageFromGoImage(goImage))
	}

	// fmt.Println("imgPath", imgPath, len(it.images))
	// it.Print()

	return it
}

func (i *imageTable) LogBoxes() {
	for i, box := range i.boxes {
		fmt.Printf("subimage %d pixels range: (xmin=%d, ymin=%d, xmax=%d, ymax=%d)\n",
			i+1, box.Xmin, box.Ymin, box.Xmax, box.Ymax)
	}
}

func (i *imageTable) Images() []imagekit.Image {
	return i.images
}

func (i *imageTable) Length() int { return len(i.images) }

func (i *imageTable) Get(index int) imagekit.Image {
	return i.images[index]
}
