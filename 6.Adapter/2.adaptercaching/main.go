package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"strings"
)

type Line struct {
	X1, Y1, X2, Y2 int
}

type VectorImage struct {
	Lines []Line
}

func NewRectangle(width, height int) *VectorImage {
	width -= 1
	height -= 1
	return &VectorImage{
		[]Line{
			Line{0, 0, width, 0},
			Line{0, 0, 0, height},
			Line{width, 0, width, height},
			Line{0, height, width, height}}}
}

type Point struct {
	X, Y int
}

type RasterImage interface {
	GetPoints() []Point
}

func DrawPoints(owner RasterImage) string {
	maxX, maxY := 0, 0
	points := owner.GetPoints()
	for _, pixel := range points {
		if pixel.X > maxX {
			maxX = pixel.X
		}
		if pixel.Y > maxY {
			maxY = pixel.Y
		}
		maxX += 1
		maxY += 1

	}
	// preallocate
	data := make([][]rune, maxY)
	for i := 0; i < maxY; i++ {
		data[i] = make([]rune, maxX)
		for j := range data[i] {
			data[i][j] = ' '
		}
	}
	for _, point := range points {
		data[point.Y][point.X] = '*'
	}
	b := strings.Builder{}
	for _, line := range data {
		b.WriteString(string(line))
		b.WriteRune('\n')
	}
	return b.String()
}

// solution
type vectorToRasterAdapter struct {
	points []Point
}

func minmax(a, b int) (int, int) {
	if a < b {
		return a, b
	} else {
		return b, a
	}
}

var pointCache = map[[16]byte][]Point{}

func (a *vectorToRasterAdapter) addLine(line Line) {
	left, right := minmax(line.X1, line.X2)
	top, bottom := minmax(line.Y1, line.Y2)
	dx := right - left
	dy := line.Y2 - line.Y1
	if dx == 0 {
		for y := top; y <= bottom; y++ {
			a.points = append(a.points, Point{left, y})
		}
	} else if dy == 0 {
		for x := left; x <= right; x++ {
			a.points = append(a.points, Point{x, top})
		}
	}
	fmt.Println("generated", len(a.points), "points")
}

func (v vectorToRasterAdapter) GetPoints() []Point {
	return v.points
}

func VectorToRaster(vi *VectorImage) RasterImage {
	adapter := vectorToRasterAdapter{}

	for _, line := range vi.Lines {
		adapter.addLineCached(line)
	}

	return adapter
}

// saving memory by hash caching

func (a *vectorToRasterAdapter) addLineCached(line Line) {
	hash := func(obj interface{}) [16]byte {
		bytes, _ := json.Marshal(obj)
		return md5.Sum(bytes)
	}
	h := hash(line)
	if pts, ok := pointCache[h]; ok {
		for _, pt := range pts {
			a.points = append(a.points, pt)
		}
	}

	left, right := minmax(line.X1, line.X2)
	top, bottom := minmax(line.Y1, line.Y2)
	dx := right - left
	dy := line.Y2 - line.Y1
	if dx == 0 {
		for y := top; y <= bottom; y++ {
			a.points = append(a.points, Point{left, y})
		}
	} else if dy == 0 {
		for x := left; x <= right; x++ {
			a.points = append(a.points, Point{x, top})
		}
	}
	pointCache[h] = a.points
	fmt.Println("generated", len(a.points), "points")
}

func main() {
	rc := NewRectangle(6, 4)
	a := VectorToRaster(rc)
	_ = VectorToRaster(rc)
	fmt.Print(DrawPoints(a))

}
