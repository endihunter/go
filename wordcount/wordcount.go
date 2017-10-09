package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)

	words := strings.Fields(s)
	length := len(words)

	for i := 0; i < length; i++ {
		m[words[i]] += 1;
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
