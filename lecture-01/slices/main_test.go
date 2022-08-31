package main

import (
	"reflect"
	"testing"
)

func TestAddInt(t *testing.T) {
	t.Parallel()

	type funcArgs struct {
		s []int
		n int
	}

	tests := []struct {
		testName string
		args     funcArgs
		expected []int
	}{
		{"slice with no elements",
			funcArgs{[]int{}, 5}, []int{}},
		{"slice with one element",
			funcArgs{[]int{33}, 10}, []int{43}},
		{"slice with two elements",
			funcArgs{[]int{0, 3}, 3}, []int{3, 6}},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			if res := addInt(tt.args.s, tt.args.n); !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("res: '%v', expected: '%v'", res, tt.expected)
			}
		})
	}
}

func TestAppendInt(t *testing.T) {
	t.Parallel()

	type funcArgs struct {
		s []int
		n int
	}

	tests := []struct {
		testName string
		args     funcArgs
		expected []int
	}{
		{"slice with no elements",
			funcArgs{[]int{}, 5}, []int{5}},
		{"slice with one element",
			funcArgs{[]int{33}, 10}, []int{33, 10}},
		{"slice with two elements",
			funcArgs{[]int{0, 3}, 3}, []int{0, 3, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			if res := appendInt(tt.args.s, tt.args.n); !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("res: '%v', expected: '%v'", res, tt.expected)
			}
		})
	}
}

func TestPrependInt(t *testing.T) {
	t.Parallel()

	type funcArgs struct {
		s []int
		n int
	}

	tests := []struct {
		testName string
		args     funcArgs
		expected []int
	}{
		{"slice with no elements",
			funcArgs{[]int{}, 5}, []int{5}},
		{"slice with one element",
			funcArgs{[]int{33}, 10}, []int{10, 33}},
		{"slice with two elements",
			funcArgs{[]int{0, 3}, 3}, []int{3, 0, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			if res := prependInt(tt.args.s, tt.args.n); !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("res: '%v', expected: '%v'", res, tt.expected)
			}
		})
	}
}

func TestRemoveIntAtIndex(t *testing.T) {
	t.Parallel()

	type funcArgs struct {
		s []int
		n int
	}

	tests := []struct {
		testName           string
		args               funcArgs
		expectedSlice      []int
		expectedRemovedInt *int
	}{
		{"slice with no elements",
			funcArgs{[]int{}, 5}, []int{}, nil},
		{"index out of slice range",
			funcArgs{[]int{0, 3}, 3}, []int{0, 3}, nil},
		{"index found and element removed",
			funcArgs{[]int{10, 33, 1}, 1},
			[]int{10, 1}, getIntPointer(33)},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			if resSlice, removedInt := removeIntAtIndex(tt.args.s, tt.args.n); !reflect.DeepEqual(resSlice, tt.expectedSlice) {
				t.Errorf("resSlice: '%v', expected: '%v'", resSlice, tt.expectedSlice)
				t.Errorf("removedInt: '%v', expectedRemovedInt: '%v'", removedInt, tt.expectedRemovedInt)
			}
		})
	}
}

func TestMergeUniqueInts(t *testing.T) {
	t.Parallel()

	type funcArgs struct {
		s0, s1 []int
	}

	tests := []struct {
		testName string
		args     funcArgs
		expected []int
	}{
		{"slices with no elements",
			funcArgs{[]int{}, []int{}}, []int{}},
		{"one slice with elements and the other empty",
			funcArgs{[]int{33, 1}, []int{}}, []int{1, 33}},
		{"two slices with elements no duplicates",
			funcArgs{[]int{33, 1}, []int{44, 2}}, []int{1, 2, 33, 44}},
		{"two slices with duplicated elements",
			funcArgs{[]int{33, 1, 1, 2}, []int{44, 44, 2}}, []int{1, 2, 33, 44}},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			if res := mergeUniqueInts(tt.args.s0, tt.args.s1); !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("res: '%v', expected: '%v'", res, tt.expected)
			}
		})
	}
}

func TestRemoveSameInts(t *testing.T) {
	t.Parallel()

	type funcArgs struct {
		s0, s1 []int
	}

	tests := []struct {
		testName string
		args     funcArgs
		expected []int
	}{
		{"slices with no elements",
			funcArgs{[]int{}, []int{}}, []int{}},
		{"one slice with elements and the other empty",
			funcArgs{[]int{33, 1}, []int{}}, []int{33, 1}},
		{"two slices with elements in common",
			funcArgs{[]int{33, 1, 0, 5, 66}, []int{5, 1}}, []int{33, 0, 66}},
		{"two slices with no elements in common",
			funcArgs{[]int{33, 1, 0, 5, 66}, []int{2, 4, 6, 8}}, []int{33, 1, 0, 5, 66}},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			if res := removeSameInts(tt.args.s0, tt.args.s1); !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("res: '%v', expected: '%v'", res, tt.expected)
			}
		})
	}
}

func TestShiftInts(t *testing.T) {
	t.Parallel()

	type funcArgs struct {
		s   []int
		i   int
		dir bool
	}

	tests := []struct {
		testName string
		args     funcArgs
		expected []int
	}{
		{"slice with no elements",
			funcArgs{[]int{}, 5, true}, []int{}},
		{"shift right once",
			funcArgs{[]int{33}, 1, true}, []int{33}},
		{"shift twice left",
			funcArgs{[]int{33, 1, 0, 5, 66}, 2, false}, []int{0, 5, 66, 33, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			if res := shiftInts(tt.args.s, tt.args.i, tt.args.dir); !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("res: '%v', expected: '%v'", res, tt.expected)
			}
		})
	}
}

func TestCopyIntSlice(t *testing.T) {
	t.Parallel()

	type funcArgs struct {
		s []int
	}

	tests := []struct {
		testName string
		args     funcArgs
		expected []int
	}{
		{"slice with no elements",
			funcArgs{[]int{}}, []int{}},
		{"slice with elements",
			funcArgs{[]int{33, 44, 0, 2}}, []int{33, 44, 0, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			if res := copyIntSlice(tt.args.s); !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("res: '%v', expected: '%v'", res, tt.expected)
			}
		})
	}
}

func TestSwapIntsByIndices(t *testing.T) {
	t.Parallel()

	type funcArgs struct {
		s []int
	}

	tests := []struct {
		testName string
		args     funcArgs
		expected []int
	}{
		{"slice with no elements",
			funcArgs{[]int{}}, []int{}},
		{"slice with one element",
			funcArgs{[]int{2}}, []int{2}},
		{"slice with even elements",
			funcArgs{[]int{33, 44, 0, 2}}, []int{44, 33, 2, 0}},
		{"slice with odd elements",
			funcArgs{[]int{33, 44, 0, 2, 5}}, []int{44, 33, 2, 0, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			if res := swapIntsByIndices(tt.args.s); !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("res: '%v', expected: '%v'", res, tt.expected)
			}
		})
	}
}

func TestSortSlice_Int(t *testing.T) {
	t.Parallel()

	type funcArgs struct {
		s       []int
		reverse bool
	}

	tests := []struct {
		testName string
		args     funcArgs
		expected []int
	}{
		{"slice with no elements",
			funcArgs{[]int{}, false}, []int{}},
		{"slice with one element",
			funcArgs{[]int{2}, false}, []int{2}},
		{"slice with many elements",
			funcArgs{[]int{33, 44, 0, 2}, false}, []int{0, 2, 33, 44}},
		{"slice with many elements reverse",
			funcArgs{[]int{33, 44, 0, 2, 5}, true}, []int{44, 33, 5, 2, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			if res := sortSlice(tt.args.s, tt.args.reverse); !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("res: '%v', expected: '%v'", res, tt.expected)
			}
		})
	}
}

func TestSortSlice_String(t *testing.T) {
	t.Parallel()

	type funcArgs struct {
		s       []string
		reverse bool
	}

	tests := []struct {
		testName string
		args     funcArgs
		expected []string
	}{
		{"slice with no elements",
			funcArgs{[]string{""}, false}, []string{""}},
		{"slice with one element",
			funcArgs{[]string{"slice"}, false}, []string{"slice"}},
		{"slice with many elements",
			funcArgs{[]string{"slice", "with", "many", "elements"}, false},
			[]string{"elements", "many", "slice", "with"}},
		{"slice with many elements reverse",
			funcArgs{[]string{"slice", "with", "many", "elements", "reverse"}, true},
			[]string{"with", "slice", "reverse", "many", "elements"}},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			if res := sortSlice(tt.args.s, tt.args.reverse); !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("res: '%v', expected: '%v'", res, tt.expected)
			}
		})
	}
}

func getIntPointer(val int) *int {
	return &val
}
