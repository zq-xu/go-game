package image

import (
	"fmt"
	"image"
	"image/color"
	_ "image/png"

	"github.com/rotisserie/eris"
)

type imageTable struct {
	boxes []BBox

	images []image.Image
}

func (i *imageTable) Print() {
	for i, box := range i.boxes {
		fmt.Printf("subimage %d pixels range: (xmin=%d, ymin=%d, xmax=%d, ymax=%d)\n",
			i+1, box.Xmin, box.Ymin, box.Xmax, box.Ymax)
	}
}

func (i *imageTable) Images() []image.Image {
	return i.images
}

func NewDungeonImageTable(imgPath string) (*imageTable, error) {
	img, err := NewDungeonImage(imgPath)
	if err != nil {
		return nil, eris.Wrapf(err, "failed to load image file %s.", imgPath)
	}

	it := &imageTable{}

	it.boxes = findBoundingBoxes(img.GoImage())

	for _, v := range it.boxes {
		it.images = append(it.images, img.Image().SubImage(image.Rect(v.Xmin, v.Ymin, v.Xmax, v.Ymax)))
	}

	// fmt.Println("imgPath", imgPath, len(it.images))
	// it.Print()

	return it, nil
}

type BBox struct {
	Xmin, Ymin, Xmax, Ymax int
}

var (
	minArea = 10
	minW    = 2
	minH    = 2
)

// findBoundingBoxes: find subimages in a images
// minArea: the area min threshold
// minW, minH: bounding box min width and height
func findBoundingBoxes(img image.Image) []BBox {
	b := img.Bounds()
	w, h := b.Dx(), b.Dy()

	visited := make([][]bool, h)
	for i := range visited {
		visited[i] = make([]bool, w)
	}

	var boxes []BBox

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if visited[y][x] {
				continue
			}

			if !isForeground(img.At(b.Min.X+x, b.Min.Y+y)) {
				visited[y][x] = true
				continue
			}

			// flood-fill using stack (8-connectivity)
			stack := []image.Point{{X: x, Y: y}}
			visited[y][x] = true

			box := BBox{Xmin: w, Ymin: h, Xmax: 0, Ymax: 0}
			area := 0

			for len(stack) > 0 {
				p := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				px, py := p.X, p.Y

				area++
				if px < box.Xmin {
					box.Xmin = px
				}
				if py < box.Ymin {
					box.Ymin = py
				}
				if px > box.Xmax {
					box.Xmax = px
				}
				if py > box.Ymax {
					box.Ymax = py
				}

				// 8 neighbors
				for dy := -1; dy <= 1; dy++ {
					for dx := -1; dx <= 1; dx++ {
						if dx == 0 && dy == 0 {
							continue
						}
						nx := px + dx
						ny := py + dy
						if nx < 0 || ny < 0 || nx >= w || ny >= h {
							continue
						}
						if visited[ny][nx] {
							continue
						}
						visited[ny][nx] = true
						if isForeground(img.At(b.Min.X+nx, b.Min.Y+ny)) {
							stack = append(stack, image.Point{X: nx, Y: ny})
						}
					}
				}
			} // end flood-fill

			width := box.Xmax - box.Xmin + 1
			height := box.Ymax - box.Ymin + 1

			if area >= minArea && width >= minW && height >= minH {
				boxes = append(boxes, box)
			}
			// otherwise, drop as noise
		}
	}
	return boxes
}

const grayThreshold = uint32(60000)

func isForeground(c color.Color) bool {
	r, g, b, a := c.RGBA() // return 0..65535
	if a == 0 {
		return false
	}
	gray := (r + g + b) / 3
	return gray < grayThreshold
}
