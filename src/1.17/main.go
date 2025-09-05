package main

import "fmt"

func binarySearch(arr []int, left, right, x int) int {
	if left > right {
		return -1
	}

	mid := (left + right) / 2

	switch {
	case x > arr[mid]:
		return binarySearch(arr, mid+1, right, x)
	case x < arr[mid]:
		return binarySearch(arr, left, mid-1, x)
	default:
		return mid

	}
}

func searchBi(arr []int, x int) int {
	l, r := 0, len(arr)-1
	idx := binarySearch(arr, l, r, x)
	return idx
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	x := 2
	fmt.Println(searchBi(arr, x))
}
