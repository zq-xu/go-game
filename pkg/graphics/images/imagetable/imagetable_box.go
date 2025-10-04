package imagetable

import (
	"image"
	"image/color"
)

const grayThreshold = uint32(60000)

var (
	minArea = 10
	minW    = 2
	minH    = 2
)

type BoundingBox interface {
	Xmin() int
	Xmax() int

	Ymin() int
	Ymax() int

	Width() int
	Height() int
}

type boundingBox struct {
	xMin, yMin, xMax, yMax int
	width, height          int
}

// NewBoundingBoxes: find subimages in a images
// minArea: the area min threshold
// minW, minH: bounding box min width and height
func NewBoundingBoxes(img image.Image) []BoundingBox {
	b := img.Bounds()
	w, h := b.Dx(), b.Dy()

	visited := make([][]bool, h)
	for i := range visited {
		visited[i] = make([]bool, w)
	}

	var boxes []BoundingBox

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

			box := &boundingBox{xMin: w, yMin: h, xMax: 0, yMax: 0}
			area := 0

			for len(stack) > 0 {
				p := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				px, py := p.X, p.Y

				area++
				if px < box.xMin {
					box.xMin = px
				}
				if py < box.yMin {
					box.yMin = py
				}
				if px > box.xMax {
					box.xMax = px
				}
				if py > box.yMax {
					box.yMax = py
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

			box.width = box.xMax - box.xMin + 1
			box.height = box.yMax - box.yMin + 1

			if area >= minArea && box.width >= minW && box.height >= minH {
				boxes = append(boxes, box)
			}
			// otherwise, drop as noise
		}
	}
	return boxes
}

func isForeground(c color.Color) bool {
	r, g, b, a := c.RGBA() // return 0..65535
	if a == 0 {
		return false
	}
	gray := (r + g + b) / 3
	return gray < grayThreshold
}

func (bb *boundingBox) Xmin() int   { return bb.xMin }
func (bb *boundingBox) Xmax() int   { return bb.xMax }
func (bb *boundingBox) Ymin() int   { return bb.yMin }
func (bb *boundingBox) Ymax() int   { return bb.yMax }
func (bb *boundingBox) Width() int  { return bb.width }
func (bb *boundingBox) Height() int { return bb.height }
