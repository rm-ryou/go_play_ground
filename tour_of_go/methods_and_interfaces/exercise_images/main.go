package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	X, Y int
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) At(x, y int) color.Color {
	v := uint8((x + y) / 2)
	return color.RGBA{
		R: v,
		G: 0,
		B: v,
		A: 255,
	}
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.X, img.Y)
}

func main() {
	m := Image{
		X: 255,
		Y: 255,
	}
	pic.ShowImage(m)
}
