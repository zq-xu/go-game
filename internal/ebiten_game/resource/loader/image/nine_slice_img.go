package image

import (
	"log"
	"sync"

	"github.com/ebitenui/ebitenui/image"
	"github.com/zq-xu/go-game/pkg/graphics"
)

type NineSliceImgLoader struct {
	lock sync.RWMutex

	nineSliceImgSet map[string]*image.NineSlice
}

func NewNineSliceImgLoader() *NineSliceImgLoader {
	return &NineSliceImgLoader{
		nineSliceImgSet: make(map[string]*image.NineSlice, 0),
	}
}

func (il *NineSliceImgLoader) MustGetNineSliceSimpleImage(path string, borderWidthHeight, centerWidthHeight int) *image.NineSlice {
	img, err := il.GetNineSliceSimpleImage(path, borderWidthHeight, centerWidthHeight)
	if err != nil {
		log.Fatal(err)
	}

	return img
}

func (il *NineSliceImgLoader) GetNineSliceSimpleImage(path string, borderWidthHeight, centerWidthHeight int) (*image.NineSlice, error) {
	img := il.getNineSliceImage(path)
	if img != nil {
		return img, nil
	}

	return il.loadNineSliceSimpleImage(path, borderWidthHeight, centerWidthHeight)
}

func (il *NineSliceImgLoader) MustGetNineSliceImage(path string, borderWidthHeight, centerWidthHeight int) *image.NineSlice {
	img, err := il.GetNineSliceImage(path, borderWidthHeight, centerWidthHeight)
	if err != nil {
		log.Fatal(err)
	}

	return img
}

func (il *NineSliceImgLoader) GetNineSliceImage(path string, centerWidth int, centerHeight int) (*image.NineSlice, error) {
	img := il.getNineSliceImage(path)
	if img != nil {
		return img, nil
	}

	return il.loadNineSliceImage(path, centerWidth, centerHeight)
}

func (il *NineSliceImgLoader) getNineSliceImage(path string) *image.NineSlice {
	il.lock.RLock()
	defer il.lock.RUnlock()

	return il.nineSliceImgSet[path]
}

func (il *NineSliceImgLoader) loadNineSliceSimpleImage(path string, borderWidthHeight, centerWidthHeight int) (*image.NineSlice, error) {
	il.lock.Lock()
	defer il.lock.Unlock()

	img, err := loadNineSliceSimpleImage(path, borderWidthHeight, centerWidthHeight)
	if err != nil {
		return nil, err
	}

	il.nineSliceImgSet[path] = img
	return img, nil
}

func (il *NineSliceImgLoader) loadNineSliceImage(path string, centerWidth int, centerHeight int) (*image.NineSlice, error) {
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
	i, err := graphics.NewImageFromFile(path)
	if err != nil {
		return nil, err
	}

	w := i.Image.Bounds().Dx()
	h := i.Image.Bounds().Dy()
	return image.NewNineSlice(i.Image,
			[3]int{(w - centerWidth) / 2, centerWidth, w - (w-centerWidth)/2 - centerWidth},
			[3]int{(h - centerHeight) / 2, centerHeight, h - (h-centerHeight)/2 - centerHeight}),
		nil
}

func loadNineSliceSimpleImage(path string, borderWidthHeight, centerWidthHeight int) (*image.NineSlice, error) {
	i, err := graphics.NewImageFromFile(path)
	if err != nil {
		return nil, err
	}

	return image.NewNineSliceSimple(i.Image, borderWidthHeight, centerWidthHeight), nil
}
