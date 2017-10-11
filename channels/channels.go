package main

import (
	"fmt"
)

func sum(data []int, c chan int) {
	sum := 0
	for _, v := range data {
		sum += v
	}

	c <- sum
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1

	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}

	close(c)
}

func main() {
	s := []int{1, 3, 5, 6, 4, 5}
	c := make(chan int)

	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)

	////////////////// Fibonacci using Channels

	ch := make(chan int, 10)
	go fibonacci(cap(ch), ch)

	for v := range ch {
		fmt.Println(v)
	}
}
