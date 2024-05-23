package main

import "fmt"

type Image interface {
	Draw()
}

type Bitmap struct {
	filename string
}

func NewBitmap(filename string) *Bitmap {
	fmt.Println("Loading image from", filename)
	return &Bitmap{filename: filename}
}

func (b *Bitmap) Draw() {
	fmt.Println("Drawing image", b.filename)
}

func DrawImage(image Image) {
	fmt.Println("About to draw the image")
	image.Draw()
	fmt.Println("Done drawing the image")
}

type LazyBitmap struct {
	filename string
	bitmap   *Bitmap
}

func (l *LazyBitmap) Draw() {
	if l.bitmap == nil {
		l.bitmap = NewBitmap(l.filename)
	}
	l.bitmap.Draw()
}

func NewLazyBitmap(filename string) *LazyBitmap {
	return &LazyBitmap{filename: filename}
}

func main() {
	bmp := NewBitmap("demo.png")
	DrawImage(bmp)

	// what if we don't draw the image
	// > Loading image from demo.png
	_ = NewBitmap("demo.png")

	fmt.Println("-----------")

	bmp2 := NewLazyBitmap("demo.png")
	DrawImage(bmp2)
	DrawImage(bmp2)
	// > Loading image starts only when DrawImage called
	// Loading happens only once even if draw image called twice
}
