package main

import (
	"fmt"
	"math"
	"strings"
)

type Abser interface {
	Abs() float64
}

type Printable interface {
	Print() string
}

type User struct {
	Name  string
	Email string
}

func (t User) Print() string {
	return t.Name
}

type IPAddr [4]byte

func (a IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", a[0], a[1], a[2], a[3])
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

func (p Point) String() string {
	return fmt.Sprintf("[%v,%v]", p.x, p.y)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	case Printable:
		fmt.Printf("Hello \"%v\"\n", v.Print())
	default:
		fmt.Printf("I don't know about type %T\n", v)
	}
}

type NegativeNumber float64

func (n NegativeNumber) Error() string {
	return fmt.Sprintf("Can not Process negative numbers: %v", float64(n))
}

func Sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0, NegativeNumber(n)
	}

	return math.Sqrt(n), nil;
}

func delimiter() {
	fmt.Println(strings.Repeat("=", 60))
}

func section(title interface{}) {
	fmt.Println()
	delimiter()
	fmt.Println(title)
	delimiter()
	fmt.Println()
}

func main() {
	section("Interfaces")
	var a Abser

	f := MyFloat(10)
	p := Point{10, 20}

	a = f
	fmt.Println(a.Abs())

	a = &p // & - because func receives a *Pointer
	fmt.Println(a.Abs())

	section("Interfaces are implemented implicitly")

	var i Printable = User{"John Doe", "john@doe.com"}
	fmt.Println("My name is:", i.Print())

	fmt.Printf("%+v, %T\n", i, i);

	section("Type Assertions")

	if _, ok := i.(Printable); ok {
		fmt.Printf("Implmenting Right Interface %T\n", i);
	}

	section("Type Switches")

	do(21)
	do("hello")
	do(User{"Endi", "endi@terranet.md"})
	do(true)

	section("Stringers - Self-Describing Interfaces")

	c := Point{100, 200}
	fmt.Println("My coordinates are:", c.String());

	section("Stringers Exercise")

	hosts := map[string]IPAddr{
		"localhost": {127, 0, 0, 1},
		"google":    {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

	section("Errors")

	numbers := []float64{4, -4}

	for _, n := range numbers {
		if val, err := Sqrt(n); err == nil {
			fmt.Printf("The SQRT(%v) is: %v\n", n, val)
		} else {
			fmt.Println(err.Error())
		}
	}
}
