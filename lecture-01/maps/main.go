package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

func main() {
	// exercise 1
	s0 := "I am actually thinking whether I am allowed to use strings package or not . " +
		"In the specs of the exercise, it is stated: << Try to use only standard golang packages >> , " +
		"but apparently strings is in the standard package , as stated here: https://pkg.go.dev/std . " +
		"So - yeah - I am going to use it. Maybe I am not allowed tho :)"

	fmt.Printf("\nCounting occurrences of string: \n'%s'\n", s0)
	_ = countWordsOccurrences(s0)

	// exercise 2
	rand.Seed(time.Now().Unix())

	s1 := generateRandomIntSlice()
	fmt.Printf("\nCounting occurrences of integers in slice: \n'%v'\n\n", s1)
	countIntsOccurrences(s1)

	// exercise 3
	s2 := generateRandomIntSlice()
	s3 := findSameInts(s1, s2)
	fmt.Printf("\nSame numbers found in slices '%v' and '%v' are:  '%v'\n\n", s1, s2, s3)

	// exercise 4
	num := rand.Intn(100)
	fibN := fibonacci(num)
	fmt.Printf("Fibonacci succession number for value '%d' is: '%d'\n\n", num, fibN)

	// exercise 5
	matrix := generateRandomMatrix(rand.Intn(10), rand.Intn(10))
	fmt.Printf("Generated random matrix before sort: '%v'\n\n", matrix)
	sortAndPrintMatrix(matrix)
}

// countWordsOccurrences counts the occurrences of each word in a given text and returns the result
func countWordsOccurrences(s string) map[string]int {
	words := strings.Fields(s)
	m := make(map[string]int)

	for _, word := range words {
		m[word] += 1
	}

	for i, v := range m {
		fmt.Println(i, "--->", v)
	}

	return m
}

// generateRandomIntSlice returns a slice of integers with random length and elements
func generateRandomIntSlice() []int {
	s := make([]int, rand.Intn(10))
	for i := 0; i < len(s); i++ {
		s[i] = rand.Intn(10)
	}

	return s
}

// countIntsOccurrences counts the occurrences of each int in a given slice of integers and returns it
func countIntsOccurrences(s []int) map[int]int {
	m := make(map[int]int)
	for _, v := range s {
		m[v] = m[v] + 1
	}

	fmt.Println(m)

	return m
}

// findSameInts finds the integers which appear in many slices and return a slice containing them
func findSameInts(slices ...[]int) []int {
	m := map[int]int{}

	var s []int

	for _, slice := range slices {
		for _, v := range slice {
			m[v] += 1
		}
	}

	for k, v := range m {
		if v > 1 {
			s = append(s, k)
		}
	}

	if len(s) == 0 {
		return []int{}
	}

	sort.Ints(s)

	return s
}

// fibonacci function implementation using memoization
func fibonacci(n int) int {
	m := make(map[int]int)
	m[0], m[1] = 0, 1

	return fib(n, m)
}

// fib checks if the number is in the cache
// if not, it calculates it, saves it into the cache, and returns it
func fib(n int, m map[int]int) int {
	i, ok := m[n]

	if ok {
		return i
	}

	m[n] = fib(n-1, m) + fib(n-2, m)

	return m[n]
}

// generateRandomMatrix generates a map of []bytes given its dimensions
func generateRandomMatrix(m, n int) map[int]map[int][]byte {
	ret := make(map[int]map[int][]byte)

	for i := 0; i < m; i++ {
		i = rand.Intn(10)
		ret[i] = make(map[int][]byte)

		for j := 0; j < n; j++ {
			j = rand.Intn(10)
			b := make([]byte, 20)
			rand.Read(b)
			ret[i][j] = b
		}
	}

	return ret
}

// sortAndPrintMatrix prints the matrix as sorted data with descending sorting by iKeys and ascending by jKeysß
func sortAndPrintMatrix(m map[int]map[int][]byte) {
	iKeys := make([]int, 0, len(m))
	for i := range m {
		iKeys = append(iKeys, i)
	}

	// could have used the method implemented for slices but wanted to try the sort std package one
	sort.Sort(sort.Reverse(sort.IntSlice(iKeys)))

	for _, i := range iKeys {
		jKeys := make([]int, 0, len(m[i]))
		for j := range m[i] {
			jKeys = append(jKeys, j)
		}

		sort.Sort(sort.Reverse(sort.IntSlice(jKeys)))

		for _, j := range jKeys {
			fmt.Printf("iKey:'%d', jKey'%d'", i, j)
		}
	}
}
