package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	length := len(words)
	m := make(map[string]int, length)

	for _, w := range words {
		m[w] += 1
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
