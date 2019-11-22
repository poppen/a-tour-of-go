package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// Image is struct
type Image struct{}

// Bounds is ...
func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 256, 256)
}

// ColorModel is ...
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

// At is ..
func (i Image) At(x, y int) color.Color {
	v := uint8(x * y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
