package imageloader

import (
	"embed"
	"sync"

	"github.com/ebitenui/ebitenui/image"

	"github.com/zq-xu/go-game/pkg/graphics/images/imagekit"
)

type NineSliceImgLoader interface {
	GetNineSliceSimpleImage(path string, borderWidthHeight, centerWidthHeight int) (*image.NineSlice, error)
	GetFixedNineSlice(path string) (*image.NineSlice, error)
	GetNineSliceImage(path string, centerWidth int, centerHeight int) (*image.NineSlice, error)
}

type nineSliceImgLoader struct {
	lock sync.RWMutex

	embedFS *embed.FS

	nineSliceImgSet map[string]*image.NineSlice
}

func NewNineSliceImgLoader(embedFS *embed.FS) NineSliceImgLoader {
	return &nineSliceImgLoader{
		embedFS:         embedFS,
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

func (il *nineSliceImgLoader) GetFixedNineSlice(path string) (*image.NineSlice, error) {
	img := il.getNineSliceImage(path)
	if img != nil {
		return img, nil
	}

	return il.loadFixedNineSliceImage(path)
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

	img, err := il.newNineSliceSimpleImage(path, borderWidthHeight, centerWidthHeight)
	if err != nil {
		return nil, err
	}

	il.nineSliceImgSet[path] = img
	return img, nil
}

func (il *nineSliceImgLoader) loadFixedNineSliceImage(path string) (*image.NineSlice, error) {
	il.lock.Lock()
	defer il.lock.Unlock()

	img, err := il.newFixedNineSliceImage(path)
	if err != nil {
		return nil, err
	}

	il.nineSliceImgSet[path] = img
	return img, nil
}

func (il *nineSliceImgLoader) loadNineSliceImage(path string, centerWidth int, centerHeight int) (*image.NineSlice, error) {
	il.lock.Lock()
	defer il.lock.Unlock()

	img, err := il.newNineSliceImage(path, centerWidth, centerHeight)
	if err != nil {
		return nil, err
	}

	il.nineSliceImgSet[path] = img
	return img, nil
}

func (il *nineSliceImgLoader) newFixedNineSliceImage(path string) (*image.NineSlice, error) {
	i, err := imagekit.NewImageFromEmbed(il.embedFS, path)
	if err != nil {
		return nil, err
	}

	return image.NewFixedNineSlice(i.Image()), nil
}

func (il *nineSliceImgLoader) newNineSliceImage(path string, centerWidth int, centerHeight int) (*image.NineSlice, error) {
	i, err := imagekit.NewImageFromEmbed(il.embedFS, path)
	if err != nil {
		return nil, err
	}

	w := i.Image().Bounds().Dx()
	h := i.Image().Bounds().Dy()
	return image.NewNineSlice(i.Image(),
			[3]int{(w - centerWidth) / 2, centerWidth, w - (w-centerWidth)/2 - centerWidth},
			[3]int{(h - centerHeight) / 2, centerHeight, h - (h-centerHeight)/2 - centerHeight}),
		nil
}

func (il *nineSliceImgLoader) newNineSliceSimpleImage(path string, borderWidthHeight, centerWidthHeight int) (*image.NineSlice, error) {
	i, err := imagekit.NewImageFromEmbed(il.embedFS, path)
	if err != nil {
		return nil, err
	}

	return image.NewNineSliceSimple(i.Image(), borderWidthHeight, centerWidthHeight), nil
}
