package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	Width, Height int
}

// ColorModel returns color.RGBAModel
func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds returns an image.Rectangle
func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.Width, img.Height)
}

// At returns the color. Using the same value as the previous img exercise (uint8(x*x + y*y + 33))
func (img Image) At(x, y int) color.Color {
	return color.RGBA{R: uint8(x*x + y*y + 33), B: 255, A: 255}
}

func main() {
	m := Image{100, 100}
	pic.ShowImage(m)
}
