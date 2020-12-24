package main

import "testing"

func IntSliceEqual(left, right []int) bool {
	if len(left) != len(right) {
		return false
	}
	for i := 0; i < len(left); i++ {
		if left[i] != right[i] {
			return false
		}
	}
	return true
}

func TestTraverseSpiral2D(t *testing.T) {
	testCase := []struct {
		Name     string
		M        [][]int
		Expected []int
		HasError bool
	}{
		{
			Name:     "Correct. 1 x 1",
			M:        [][]int{{1}},
			Expected: []int{1},
			HasError: false,
		},
		{
			Name:     "Correct. 2 x 2",
			M:        [][]int{{1, 2}, {4, 3}},
			Expected: []int{1, 2, 3, 4},
			HasError: false,
		},
		{
			Name:     "Correct. 3 x 3",
			M:        [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}},
			Expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			HasError: false,
		},
		{
			Name:     "Correct. 4 x 4",
			M:        [][]int{{1, 2, 3, 4}, {12, 13, 14, 5}, {11, 16, 15, 6}, {10, 9, 8, 7}},
			Expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			HasError: false,
		},
		{
			Name:     "Correct. 5 x 5",
			M:        [][]int{{1, 2, 3, 4, 5}, {16, 17, 18, 19, 6}, {15, 24, 25, 20, 7}, {14, 23, 22, 21, 8}, {13, 12, 11, 10, 9}},
			Expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25},
			HasError: false,
		},
		{
			Name:     "Error. 0 x 0",
			M:        [][]int{{}},
			Expected: []int{},
			HasError: true,
		},
		{
			Name:     "Error. 1 x 2",
			M:        [][]int{{4, 3}},
			Expected: []int{},
			HasError: true,
		},
		{
			Name:     "Error. 2 x 1",
			M:        [][]int{{1}, {2}},
			Expected: []int{},
			HasError: true,
		},
	}
	for _, c := range testCase {
		t.Run(c.Name, func(t *testing.T) {
			flat, err := TraverseSpiral2D(c.M)
			if c.HasError && err != nil {
				return
			}
			if err != nil {
				t.Error("unexpected error", err)
			}
			if !IntSliceEqual(c.Expected, flat) {
				t.Error("expected", c.Expected, "but got", flat)
			}
		})
	}
}
