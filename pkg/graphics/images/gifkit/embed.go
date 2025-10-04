package gifkit

import (
	"embed"

	"github.com/rotisserie/eris"
	"github.com/zq-xu/go-game/pkg/graphics/images/imagekit"
)

type embedGif struct {
	Gif
}

func NewEmbedGif(embedFS *embed.FS, imgPaths []string) (Gif, error) {
	images := make([]imagekit.Image, 0)

	for _, v := range imgPaths {
		img, err := imagekit.NewImageFromEmbed(embedFS, v)
		if err != nil {
			return nil, eris.Wrap(err, "new image from embed failed")
		}
		images = append(images, img)

	}

	return &embedGif{Gif: NewGIF(images...)}, nil
}
