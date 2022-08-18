package main

import "fmt"

// fibonacci returns a closure that returns successive Fibonacci numbers
func fibonacci() func() int {
	first, second := 0, 1

	return func() int {
		res := first
		first, second = second, res+second

		return res
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 33; i++ {
		fmt.Println(f())
	}
}
