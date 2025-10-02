package imageloader

import (
	"embed"
	"sync"

	"github.com/rotisserie/eris"
	"github.com/zq-xu/go-game/pkg/graphics/images/imagekit"
)

type ImageLoader interface {
	GetImage(path string) (imagekit.Image, error)
}

type imgLoader struct {
	lock sync.RWMutex

	embedFS *embed.FS

	imgSet map[string]imagekit.Image
}

func NewimgLoader(embedFS *embed.FS) ImageLoader {
	return &imgLoader{
		embedFS: embedFS,
		imgSet:  make(map[string]imagekit.Image, 0),
	}
}

func (il *imgLoader) GetImage(path string) (imagekit.Image, error) {
	img := il.getImage(path)
	if img != nil {
		return img, nil
	}

	return il.loadImage(path)
}

func (il *imgLoader) getImage(path string) imagekit.Image {
	il.lock.RLock()
	defer il.lock.RUnlock()

	return il.imgSet[path]
}

func (il *imgLoader) loadImage(path string) (imagekit.Image, error) {
	il.lock.Lock()
	defer il.lock.Unlock()

	img, err := il.newImageFromEmbed(path)
	if err != nil {
		return nil, eris.Wrap(err, "new image from embed failed.")
	}

	il.imgSet[path] = img
	return img, nil
}

func (il *imgLoader) newImageFromEmbed(path string) (imagekit.Image, error) {
	return imagekit.NewImageFromEmbed(il.embedFS, path)
}
