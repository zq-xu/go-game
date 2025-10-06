package imagetable

import (
	"embed"
	"image"

	"github.com/rotisserie/eris"
	"github.com/zq-xu/go-game/pkg/graphics/images/imagekit"
	"github.com/zq-xu/go-game/pkg/logs"
)

type ImageTable interface {
	Images() []imagekit.Image
	Length() int

	// for debug
	LogBoxes()
}

type imageTable struct {
	name   string
	boxes  []BoundingBox
	images []imagekit.Image
}

// NewDungeonImageTable
func NewImageTable(name string, img imagekit.Image) ImageTable {
	it := &imageTable{
		name:  name,
		boxes: NewBoundingBoxes(img.GoImage()),
	}

	for _, v := range it.boxes {
		goImage := img.Image().SubImage(image.Rect(v.Xmin(), v.Ymin(), v.Xmax(), v.Ymax()))
		it.images = append(it.images, imagekit.NewBasicImageFromGoImage(goImage))
	}

	return it
}

func NewImageTableFromEmbed(embedFS *embed.FS, name, imgPath string) (ImageTable, error) {
	img, err := imagekit.NewImageFromEmbed(embedFS, imgPath)
	if err != nil {
		return nil, eris.Wrapf(err, "failed to load image file %s.", img)
	}

	return NewImageTable(name, img), nil
}

func (i *imageTable) LogBoxes() {
	logs.Logger.Debugf("imagebox %s:", i.name)
	for i, box := range i.boxes {
		logs.Logger.Debugf("  subimage %d pixels range: (xmin=%d, ymin=%d, xmax=%d, ymax=%d, width=%d, height=%d)\n",
			i+1, box.Xmin(), box.Ymin(), box.Xmax(), box.Ymax(), box.Width(), box.Height())
	}
}

func (i *imageTable) Images() []imagekit.Image {
	return i.images
}

func (i *imageTable) Length() int { return len(i.images) }

func (i *imageTable) Get(index int) imagekit.Image {
	return i.images[index]
}
