package image

import (
	"log"
	"sync"

	"github.com/zq-xu/go-game/pkg/graphics"
)

type ImgLoader struct {
	lock sync.RWMutex

	imgSet map[string]*graphics.Image
}

func NewImgLoader() *ImgLoader {
	return &ImgLoader{
		imgSet: make(map[string]*graphics.Image, 0),
	}
}

func (il *ImgLoader) MustGetImage(path string) *graphics.Image {
	img, err := il.GetImage(path)
	if err != nil {
		log.Fatal(err)
	}

	return img
}

func (il *ImgLoader) GetImage(path string) (*graphics.Image, error) {
	img := il.getImage(path)
	if img != nil {
		return img, nil
	}

	return il.loadImage(path)
}

func (il *ImgLoader) getImage(path string) *graphics.Image {
	il.lock.RLock()
	defer il.lock.RUnlock()

	return il.imgSet[path]
}

func (il *ImgLoader) loadImage(path string) (*graphics.Image, error) {
	il.lock.Lock()
	defer il.lock.Unlock()

	img, err := graphics.NewImageFromFile(path)
	if err != nil {
		return nil, err
	}

	il.imgSet[path] = img
	return img, nil
}
