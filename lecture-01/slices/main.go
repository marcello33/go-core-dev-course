package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	// init needed vars
	s0 := []int{0, 1, 2}

	rand.Seed(time.Now().UnixNano())

	fmt.Printf("\n\nInitial slice: '%v' \n\n", s0)

	// exercise 1
	n := rand.Intn(10)
	addInt(s0, n)
	fmt.Printf("Slice after adding '%d' to each element: '%v' \n\n", n, s0)

	// exercise 2
	n = rand.Intn(10)
	s0 = appendInt(s0, n)
	fmt.Printf("Slice after appending '%d': '%v' \n\n", n, s0)

	// exercise 3
	n = rand.Intn(10)
	s0 = prependInt(s0, n)
	fmt.Printf("Slice after prepending '%d': '%v' \n\n", n, s0)

	// exercise 4
	fmt.Printf("Slice before removing first element: '%v' \n", s0)
	s0, val := removeIntAtIndex(s0, 0)

	if val != nil {
		fmt.Printf("Slice after removing element '%d': '%v' \n\n", *val, s0)
	}

	// exercise 5
	fmt.Printf("Slice before removing last element: '%v' \n", s0)
	s0, val = removeIntAtIndex(s0, len(s0)-1)

	if val != nil {
		fmt.Printf("Slice after removing element '%d': '%v' \n\n", *val, s0)
	}

	// exercise 6
	n = rand.Intn(len(s0))

	fmt.Printf("Slice before removing element at index '%d': '%v' \n", n, s0)
	s0, val = removeIntAtIndex(s0, n)

	if val != nil {
		fmt.Printf("Slice after removing element '%d': '%v' \n\n", *val, s0)
	}

	// exercise 7
	s1 := []int{11, 22, 33}
	s2 := mergeUniqueInts(s0, s1)
	fmt.Printf("Merged '%v' and '%v' : result is '%v' \n\n", s0, s1, s2)

	// exercise 8
	s3 := []int{46, 11, 13}
	s4 := removeSameInts(s2, s3)
	fmt.Printf("Removed elements of slice '%v' from '%v'. Result: '%v'\n\n", s3, s2, s4)

	// exercise 9, 10, 11, 12
	n = rand.Intn(len(s4))
	fmt.Printf("Slice before shifting of '%d' positions to right: '%v' \n", n, s4)
	s5 := shiftInts(s4, n, true)
	fmt.Printf("Slice after shifting: '%v'\n\n", s5)

	n = rand.Intn(len(s5))
	fmt.Printf("Slice before shifting of '%d' positions to left: '%v' \n", n, s5)
	s5 = shiftInts(s5, n, false)
	fmt.Printf("Slice after shifting: '%v'\n\n", s5)

	// exercise 13
	fmt.Printf("Slice before copy: '%v', address: '%p'\n", s5, &s5)
	s6 := copyIntSlice(s5)
	fmt.Printf("Slice after copy:  '%v', address: '%p'\n\n", s6, &s6)

	// exercise 14
	fmt.Printf("Slice before swap: '%v'\n", s6)
	s6 = swapIntsByIndices(s6)
	fmt.Printf("Slice after swap:  '%v'\n\n", s6)

	// exercise 15a
	fmt.Printf("Int slice before revert sort:  '%v'\n", s6)
	s6 = sortSlice(s6, true)
	fmt.Printf("Int slice after revert sort:   '%v'\n\n", s6)

	// exercise 15b
	s7 := []string{"pretty", "nice", "golang", "dev", "course", "at", "polygon"}
	fmt.Printf("String slice before direct sort:  '%v'\n", s7)
	s8 := sortSlice(s7, false)
	fmt.Printf("String slice after direct sort:   '%v'\n\n", s8)
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

// removeSameInts removes elements from 'dest' slice when they exist in 'source' slice
func removeSameInts(dest, source []int) []int {
	m := make(map[int]bool)

	for _, val := range source {
		m[val] = true
	}

	for i := 0; i < len(dest); i++ {
		if _, ok := m[dest[i]]; ok {
			dest, _ = removeIntAtIndex(dest, i)
			i--
		}
	}

	return dest
}

// shiftInts shifts all int elements of a given slice a given number of times
// The shift happens to right in case the bool is set to true, left otherwise
func shiftInts(s []int, i int, right bool) []int {
	if len(s) < 2 || i >= len(s) || i <= 0 {
		return s
	}

	if right {
		return append(s[len(s)-i:], s[:len(s)-i]...)
	} else {
		return append(s[i:], s[:i]...)
	}
}

// copyIntSlice returns a copy of the given int slice
func copyIntSlice(src []int) []int {
	dst := make([]int, len(src))
	copy(dst, src)

	return dst
}

// swapIntsByIndices swaps all even index elements with the nearest odd indices ones
func swapIntsByIndices(s []int) []int {
	for j := 1; j < len(s); j += 2 {
		s[j-1], s[j] = s[j], s[j-1]
	}

	return s
}

// sortSlice sorts any slice of ints or strings in direct or reverse order
func sortSlice[T int | string](s []T, reverse bool) []T {
	sort.Slice(s, func(i, j int) bool {
		if reverse {
			return s[i] > s[j]
		} else {
			return s[i] < s[j]
		}
	})

	return s
}
