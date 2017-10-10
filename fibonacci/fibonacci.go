package main

import (
	"math"
	"fmt"
)

// Example of closure:
// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func (x, y int) int {
	return func(x, y int) int {
		return x + y
	}
}

// @todo return rounded int value
func fib(n int) float64 {
	return math.Pow((math.Sqrt(5)+1)/2, float64(n)) / math.Sqrt(5)
}

func main() {
	//numbers := make([]float64, 50)

	//for i := 0; i < len(numbers); i++ {
	//	numbers[i] = fib(i);
	//}

	f := fibonacci();
	x, y := int(0), int(0)

	for j := 0; j < 1000000; j++ {
		numbers := make([]int, 10)

		for i := 0; i < len(numbers); i++ {
			if i <= 1 {
				numbers[i] = f(x, i)
			} else {
				x = numbers[i-2]
				y = numbers[i-1]

				numbers[i] = f(x, y)
			}

		}
	}
	fmt.Println("Done")
}
