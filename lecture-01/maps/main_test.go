package main

import (
	"reflect"
	"testing"
)

func TestCountWordsOccurrences(t *testing.T) {
	t.Parallel()

	type funcArgs struct {
		s string
	}

	tests := []struct {
		testName string
		args     funcArgs
		expected map[string]int
	}{
		{"empty string",
			funcArgs{""}, map[string]int{}},
		{"string with one word",
			funcArgs{"s"}, map[string]int{"s": 1}},
		{"string with many words",
			funcArgs{"yes no maybe yes maybe no ok"},
			map[string]int{"yes": 2, "no": 2, "maybe": 2, "ok": 1}},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			if res := countWordsOccurrences(tt.args.s); !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("res: '%v', expected: '%v'", res, tt.expected)
			}
		})
	}
}

func TestCountIntsOccurrences(t *testing.T) {
	t.Parallel()

	type funcArgs struct {
		s []int
	}

	tests := []struct {
		testName string
		args     funcArgs
		expected map[int]int
	}{
		{"slice with no elements",
			funcArgs{[]int{}}, map[int]int{}},
		{"slice with one element",
			funcArgs{[]int{33}}, map[int]int{33: 1}},
		{"slice with two elements and no duplicates",
			funcArgs{[]int{0, 3}}, map[int]int{0: 1, 3: 1}},
		{"slice with many elements and duplicates",
			funcArgs{[]int{0, 3, 1, 2, 2, 3}}, map[int]int{0: 1, 3: 2, 1: 1, 2: 2}},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			if res := countIntsOccurrences(tt.args.s); !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("res: '%v', expected: '%v'", res, tt.expected)
			}
		})
	}
}

func TestFindSameInts(t *testing.T) {
	t.Parallel()

	type funcArgs struct {
		s0 []int
		s1 []int
	}

	tests := []struct {
		testName string
		args     funcArgs
		expected []int
	}{
		{"slices with no elements",
			funcArgs{[]int{}, []int{}}, []int{}},
		{"one slice with elements and the other empty",
			funcArgs{[]int{}, []int{1, 3, 2}}, []int{}},
		{"two slices with elements no duplicates",
			funcArgs{[]int{33, 1}, []int{44, 2}}, []int{}},
		{"two slices with duplicated elements",
			funcArgs{[]int{33, 1, 1, 2}, []int{44, 44, 2}}, []int{1, 2, 44}},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			if res := findSameInts(tt.args.s0, tt.args.s1); !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("res: '%#v', expected: '%#v'", res, tt.expected)
			}
		})
	}
}

func TestFibonacci(t *testing.T) {
	t.Parallel()

	type funcArgs struct {
		n int
	}

	tests := []struct {
		testName string
		args     funcArgs
		expected int
	}{
		{"one",
			funcArgs{1}, 1},
		{"eight",
			funcArgs{8}, 21},
		{"thirty-three",
			funcArgs{33}, 3524578},
		{"hundred",
			funcArgs{100}, 3736710778780434371},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			if res := fibonacci(tt.args.n); !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("res: '%v', expected: '%v'", res, tt.expected)
			}
		})
	}
}
