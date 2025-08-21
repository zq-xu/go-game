package image

import (
	"sync"

	"github.com/sirupsen/logrus"
)

var globalImageLoader *imgLoader

type imgLoader struct {
	lock sync.RWMutex

	imgSet map[string]*basicImage
}

func init() {
	globalImageLoader = newimgLoader()
}

// GetImage
func GetImage(path string) *basicImage {
	img, err := globalImageLoader.GetImage(path)
	if err != nil {
		logrus.Fatalf("get image failed. %s", err)
		return nil
	}
	return img
}

func newimgLoader() *imgLoader {
	return &imgLoader{
		imgSet: make(map[string]*basicImage, 0),
	}
}

func (il *imgLoader) GetImage(path string) (*basicImage, error) {
	img := il.getImage(path)
	if img != nil {
		return img, nil
	}

	return il.loadImage(path)
}

func (il *imgLoader) getImage(path string) *basicImage {
	il.lock.RLock()
	defer il.lock.RUnlock()

	return il.imgSet[path]
}

func (il *imgLoader) loadImage(path string) (*basicImage, error) {
	il.lock.Lock()
	defer il.lock.Unlock()

	img, err := NewImageFromFile(path)
	if err != nil {
		return nil, err
	}

	il.imgSet[path] = img
	return img, nil
}
