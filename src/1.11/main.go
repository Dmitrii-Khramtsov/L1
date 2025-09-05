package main

import (
	"fmt"
	"math"
)

func inerJoin(x, y []int) []int {
	minLen := int(math.Min(float64(len(x)), float64(len(y))))
	var min, max []int
	if minLen == len(x) {
		min = x
		max = y
	} else {
		min = y
		max = x
	}

	m := make(map[int]struct{})
	for _, v := range min {
		m[v] = struct{}{}
	}

	out := make([]int, 0)
	for _, v := range max {
		if _, ok := m[v]; ok {
			out = append(out, v)
		}
	}

	return out
}

func main() {
	a := []int{1, 2, 3}
	b := []int{2, 3, 4}
	fmt.Println(inerJoin(a, b))
}
