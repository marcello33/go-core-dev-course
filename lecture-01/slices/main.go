package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// init needed vars
	s0 := []int{0, 1, 2}

	rand.Seed(time.Now().UnixNano())

	fmt.Printf("\n\nInitial slice: '%v' \n\n", s0)

	n := rand.Intn(10)
	addInt(s0, n)
	fmt.Printf("Slice after adding '%d' to each element: '%v' \n\n", n, s0)

	n = rand.Intn(10)
	s0 = appendInt(s0, n)
	fmt.Printf("Slice after appending '%d': '%v' \n\n", n, s0)

	n = rand.Intn(10)
	s0 = prependInt(s0, n)
	fmt.Printf("Slice after prepending '%d': '%v' \n\n", n, s0)

	// remove first elem
	fmt.Printf("Slice before removing first element: '%v' \n", s0)
	s0, val := removeIntAtIndex(s0, 0)

	if val != nil {
		fmt.Printf("Slice after removing element '%d': '%v' \n\n", *val, s0)
	}

	// remove last elem
	fmt.Printf("Slice before removing last element: '%v' \n", s0)
	s0, val = removeIntAtIndex(s0, len(s0)-1)

	if val != nil {
		fmt.Printf("Slice after removing element '%d': '%v' \n\n", *val, s0)
	}

	// remove elem at random index
	n = rand.Intn(len(s0))

	fmt.Printf("Slice before removing element at index '%d': '%v' \n", n, s0)
	s0, val = removeIntAtIndex(s0, n)

	if val != nil {
		fmt.Printf("Slice after removing element '%d': '%v' \n\n", *val, s0)
	}

	s1 := []int{11, 22, 33}
	s2 := mergeUniqueInts(s0, s1)
	fmt.Printf("Merged '%v' and '%v' : result is '%v' \n\n", s0, s1, s2)
}

// addInt adds n to each and every element in the slice s
func addInt(s []int, n int) {
	for i := range s {
		s[i] += n
	}
}

// appendInt appends n to the end of the slice s
func appendInt(s []int, n int) []int {
	return append(s, n)
}

// prependInt prepends n to the beginning of the slice s
func prependInt(s []int, n int) []int {
	s = append(s, 0)
	copy(s[1:], s)
	s[0] = n

	return s
}

// removeIntAtIndex removes the element at a given index from the slice, returning the element and the slice itself
func removeIntAtIndex(s []int, i int) ([]int, *int) {
	if len(s) == 0 || i < 0 || i >= len(s) {
		return s, nil
	}

	v := s[i]

	return append(s[:i], s[i+1:]...), &v
}

// mergeUniqueInts merges slices of int, removes the duplicates from the result, and returns it
func mergeUniqueInts(slices ...[]int) []int {
	uniqueMap := map[int]bool{}

	for _, slice := range slices {
		for _, v := range slice {
			uniqueMap[v] = true
		}
	}

	res := make([]int, 0, len(uniqueMap))

	for key := range uniqueMap {
		res = append(res, key)
	}

	return res
}
