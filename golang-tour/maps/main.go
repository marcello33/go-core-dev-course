package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

// WordCount counts the occurrences of a word in the provided string
func WordCount(s string) map[string]int {
	res := make(map[string]int)
	// use package strings.Fields to split the string
	for _, v := range strings.Fields(s) {
		res[v]++
	}

	return res
}

func main() {
	wc.Test(WordCount)
}
