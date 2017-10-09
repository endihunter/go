package main

import (
	"math"
	"fmt"
)

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(x, y int64) int64 {
	return func(x, y int64) int64 {
		return x + y
	}
}

func fib(n int) float64 {
	return math.Pow((math.Sqrt(5)+1)/2, float64(n)) / math.Sqrt(5)
}

func main() {
	//numbers := make([]float64, 50)

	//for i := 0; i < len(numbers); i++ {
	//	numbers[i] = fib(i);
	//}

	numbers := make([]int64, 50)
	f := fibonacci()
	x, y := int64(0), int64(0)

	for i := 0; i < len(numbers); i++ {
		if i <= 1 {
			numbers[i] = f(x, int64(i))
		} else {
			x = numbers[i-2]
			y = numbers[i-1]

			numbers[i] = f(x, y)
		}

	}
	fmt.Println(numbers)
}
