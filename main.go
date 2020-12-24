package main

import (
	"fmt"
)

func main() {
	m0 := [][]int{{1}}
	m1 := [][]int{{1, 2}, {4, 3}}
	m2 := [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}
	m3 := [][]int{{1, 2, 3, 4}, {12, 13, 14, 5}, {11, 16, 15, 6}, {10, 9, 8, 7}}
	fmt.Println(TraverseSpiral2D(m0))
	fmt.Println(TraverseSpiral2D(m1))
	fmt.Println(TraverseSpiral2D(m2))
	fmt.Println(TraverseSpiral2D(m3))
}
