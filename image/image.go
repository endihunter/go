package main

import (
	"image"
	"fmt"
	"golang.org/x/tour/pic"
	"image/color"
	"math/rand"
	"math"
)

type Image struct {
	width, height int
}

func (i *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.width, i.height)
}

func (i *Image) At(x, y int) color.Color {
	return color.RGBA{
		randUint8(x%255, 255),
		randUint8(0, y%255),
		randUint8(0, (x^y)%255),
		randUint8(200, 255),
	}
}

func randUint8(min int, max int) uint8 {
	diff := math.Max(float64(max-min), float64(1))

	return uint8(min + rand.Intn(int(diff)))
}

func main() {
	i := image.NewRGBA(image.Rect(0, 0, 100, 100))

	fmt.Println(i.Bounds())
	fmt.Println(i.At(50, 50).RGBA())

	///////////////////////

	m := Image{100, 100}
	pic.ShowImage(&m)
}
