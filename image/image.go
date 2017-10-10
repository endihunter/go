package main

import (
	"image"
	"fmt"
)

func main() {
	i := image.NewRGBA(image.Rect(0, 0, 100, 100))

	fmt.Println(i.Bounds())
	fmt.Println(i.At(50, 50).RGBA())
}
