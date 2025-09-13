package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	X, Y int
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.X, i.Y)
}

func (i Image) At(x, y int) color.Color {
	v := uint8((i.X + i.Y) / 2)
	return color.RGBA{
		R: 0,
		G: 0,
		B: v,
		A: 255,
	}

}

func main() {
	m := Image{
		X: 256,
		Y: 256,
	}
	pic.ShowImage(m)
}
