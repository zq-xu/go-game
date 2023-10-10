package loader

type Loader struct {
	FontLoader *FontLoader

	ImageLoader *ImageLoader
}

func NewLoader() *Loader {
	return &Loader{
		FontLoader:  NewFontLoader(),
		ImageLoader: NewImageLoader(),
	}
}
