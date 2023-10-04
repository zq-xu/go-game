package font

import "golang.org/x/image/font"

type Interface interface {
	Face() font.Face
	TitleFace() font.Face
	BigTitleFace() font.Face

	// close()
}

type Config struct {
	DPI float64

	BigTitleFaceSize int
	TitleFaceSize    int
	FaceSize         int
	SmallFaceSize    int
}

// func (f *fonts) close() {
// 	if f.face != nil {
// 		_ = f.face.Close()
// 	}

// 	if f.titleFace != nil {
// 		_ = f.titleFace.Close()
// 	}

// 	if f.bigTitleFace != nil {
// 		_ = f.bigTitleFace.Close()
// 	}
// }
