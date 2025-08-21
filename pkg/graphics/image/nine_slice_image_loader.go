package image

import (
	"sync"

	"github.com/ebitenui/ebitenui/image"
)

var globalNineSliceImgLoader *nineSliceImgLoader

type nineSliceImgLoader struct {
	lock sync.RWMutex

	nineSliceImgSet map[string]*image.NineSlice
}

func init() {
	globalNineSliceImgLoader = newnineSliceImgLoader()
}

// GetNineSliceSimpleImage
func GetNineSliceSimpleImage(path string, borderWidthHeight, centerWidthHeight int) (*image.NineSlice, error) {
	return globalNineSliceImgLoader.GetNineSliceSimpleImage(path, borderWidthHeight, centerWidthHeight)
}

// GetNineSliceImage
func GetNineSliceImage(path string, centerWidth int, centerHeight int) (*image.NineSlice, error) {
	return globalNineSliceImgLoader.GetNineSliceImage(path, centerWidth, centerHeight)
}

func newnineSliceImgLoader() *nineSliceImgLoader {
	return &nineSliceImgLoader{
		nineSliceImgSet: make(map[string]*image.NineSlice, 0),
	}
}

func (il *nineSliceImgLoader) GetNineSliceSimpleImage(path string, borderWidthHeight, centerWidthHeight int) (*image.NineSlice, error) {
	img := il.getNineSliceImage(path)
	if img != nil {
		return img, nil
	}

	return il.loadNineSliceSimpleImage(path, borderWidthHeight, centerWidthHeight)
}

func (il *nineSliceImgLoader) GetNineSliceImage(path string, centerWidth int, centerHeight int) (*image.NineSlice, error) {
	img := il.getNineSliceImage(path)
	if img != nil {
		return img, nil
	}

	return il.loadNineSliceImage(path, centerWidth, centerHeight)
}

func (il *nineSliceImgLoader) getNineSliceImage(path string) *image.NineSlice {
	il.lock.RLock()
	defer il.lock.RUnlock()

	return il.nineSliceImgSet[path]
}

func (il *nineSliceImgLoader) loadNineSliceSimpleImage(path string, borderWidthHeight, centerWidthHeight int) (*image.NineSlice, error) {
	il.lock.Lock()
	defer il.lock.Unlock()

	img, err := loadNineSliceSimpleImage(path, borderWidthHeight, centerWidthHeight)
	if err != nil {
		return nil, err
	}

	il.nineSliceImgSet[path] = img
	return img, nil
}

func (il *nineSliceImgLoader) loadNineSliceImage(path string, centerWidth int, centerHeight int) (*image.NineSlice, error) {
	il.lock.Lock()
	defer il.lock.Unlock()

	img, err := loadNineSliceImage(path, centerWidth, centerHeight)
	if err != nil {
		return nil, err
	}

	il.nineSliceImgSet[path] = img
	return img, nil
}

func loadNineSliceImage(path string, centerWidth int, centerHeight int) (*image.NineSlice, error) {
	i, err := NewImageFromFile(path)
	if err != nil {
		return nil, err
	}

	w := i.img.Bounds().Dx()
	h := i.img.Bounds().Dy()
	return image.NewNineSlice(i.img,
			[3]int{(w - centerWidth) / 2, centerWidth, w - (w-centerWidth)/2 - centerWidth},
			[3]int{(h - centerHeight) / 2, centerHeight, h - (h-centerHeight)/2 - centerHeight}),
		nil
}

func loadNineSliceSimpleImage(path string, borderWidthHeight, centerWidthHeight int) (*image.NineSlice, error) {
	i, err := NewImageFromFile(path)
	if err != nil {
		return nil, err
	}

	return image.NewNineSliceSimple(i.img, borderWidthHeight, centerWidthHeight), nil
}
