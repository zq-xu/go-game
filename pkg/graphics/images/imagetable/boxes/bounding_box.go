package boxes

import (
	"image"
	"image/color"
)

// ==== 参数配置 ====

var (
	MinBoxWidth   = 8    // 最小 box 宽度
	MinBoxHeight  = 8    // 最小 box 高度
	MergeGapX     = 10   // 列之间最大合并间距
	MergeGapY     = 10   // 行之间最大合并间距
	MinAlphaLevel = 5000 // 前景像素 alpha 阈值 (0-65535)
)

// ==== 核心类型 ====

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

// ==== 主入口 ====

func NewBoundingBoxes(img image.Image) []BoundingBox {
	rows := detectRanges(img, true, MergeGapY)
	cols := detectRanges(img, false, MergeGapX)

	return buildBoxes(img, rows, cols)
}

// ==== 子逻辑：检测行/列 ====

func detectRanges(img image.Image, horizontal bool, mergeGap int) []intRange {
	mask := projectForeground(img, horizontal)
	ranges := findContinuousRanges(mask, mergeGap)
	offset := 0
	if !horizontal {
		offset = img.Bounds().Min.X
	} else {
		offset = img.Bounds().Min.Y
	}
	for i := range ranges {
		ranges[i].start += offset
		ranges[i].end += offset
	}
	return ranges
}

// ==== 计算投影 ====

func projectForeground(img image.Image, horizontal bool) []bool {
	b := img.Bounds()
	w, h := b.Dx(), b.Dy()

	if horizontal {
		has := make([]bool, h)
		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {
				if isForeground(img.At(b.Min.X+x, b.Min.Y+y)) {
					has[y] = true
					break
				}
			}
		}
		return has
	}

	has := make([]bool, w)
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			if isForeground(img.At(b.Min.X+x, b.Min.Y+y)) {
				has[x] = true
				break
			}
		}
	}
	return has
}

// ==== 找连续区间并合并 ====

type intRange struct{ start, end int }

func findContinuousRanges(arr []bool, mergeGap int) []intRange {
	var raw []intRange
	inRange := false
	start := 0

	for i, val := range arr {
		switch {
		case val && !inRange:
			inRange = true
			start = i
		case !val && inRange:
			inRange = false
			raw = append(raw, intRange{start, i})
		}
	}
	if inRange {
		raw = append(raw, intRange{start, len(arr)})
	}
	return mergeCloseRanges(raw, mergeGap)
}

func mergeCloseRanges(ranges []intRange, gap int) []intRange {
	if len(ranges) == 0 {
		return nil
	}
	merged := []intRange{ranges[0]}
	for i := 1; i < len(ranges); i++ {
		last := &merged[len(merged)-1]
		cur := ranges[i]
		if cur.start-last.end <= gap {
			last.end = cur.end
		} else {
			merged = append(merged, cur)
		}
	}
	return merged
}

// ==== 构建 Box ====

func buildBoxes(img image.Image, rows, cols []intRange) []BoundingBox {
	var boxes []BoundingBox
	for _, ry := range rows {
		for _, rx := range cols {
			rect := image.Rect(rx.start, ry.start, rx.end, ry.end)
			if !isValidBox(img, rect) {
				continue
			}
			box := makeBox(rect)
			if box.Width() >= MinBoxWidth && box.Height() >= MinBoxHeight {
				boxes = append(boxes, box)
			}
		}
	}
	return boxes
}

func makeBox(r image.Rectangle) *boundingBox {
	return &boundingBox{
		xMin:   r.Min.X,
		yMin:   r.Min.Y,
		xMax:   r.Max.X - 1,
		yMax:   r.Max.Y - 1,
		width:  r.Dx(),
		height: r.Dy(),
	}
}

// ==== 判断 Box 有效性 ====

func isValidBox(img image.Image, rect image.Rectangle) bool {
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			if isForeground(img.At(x, y)) {
				return true
			}
		}
	}
	return false
}

// ==== 前景判定 ====

func isForeground(c color.Color) bool {
	_, _, _, a := c.RGBA()
	return a > uint32(MinAlphaLevel)
}

// ==== 接口实现 ====

func (bb *boundingBox) Xmin() int   { return bb.xMin }
func (bb *boundingBox) Xmax() int   { return bb.xMax }
func (bb *boundingBox) Ymin() int   { return bb.yMin }
func (bb *boundingBox) Ymax() int   { return bb.yMax }
func (bb *boundingBox) Width() int  { return bb.width }
func (bb *boundingBox) Height() int { return bb.height }
