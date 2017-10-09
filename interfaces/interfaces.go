package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type Printable interface {
	Print() string
}

type User struct {
	Name string
}

func (t User) Print() string {
	return t.Name
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f);
	}
	return float64(f);
}

type Point struct {
	x, y float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

func main() {
	var a Abser

	f := MyFloat(10)
	p := Point{10, 20}

	a = f
	fmt.Println(a.Abs())

	a = &p // & - because func receives a *Pointer
	fmt.Println(a.Abs())

	// Interfaces are implemented implicitly
	var i Printable = User{"John Doe"}
	fmt.Println("My name is:", i.Print())

	fmt.Printf("%v, %T", i, i);
}
