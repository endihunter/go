package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

func (p *Point) Sum() int {
	return p.x + p.y
}

func (p *Point) Scale(by int) {
	p.x *= by
	p.y *= by
}

func main() {
	p := Point{10, 20}

	fmt.Println("Pointer:", p);
	fmt.Println("Sum: ", p.Sum());
	p.Scale(5);
	fmt.Println("Incremented: ", p.Sum());
}
