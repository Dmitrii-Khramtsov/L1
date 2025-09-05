package main

import "fmt"

func quickSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	pivot := arr[len(arr)/2]
	left, right := 0, len(arr)-1

	for left <= right {
		for arr[left] < pivot {
			left++
		}
		for arr[right] > pivot {
			right--
		}
		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}
	if right > 0 {
		quickSort(arr[:right+1])
	}
	if left < len(arr) {
		quickSort(arr[left:])
	}
}

func main() {
	arr := []int{3, 7, 2, 5, 1, 9, 6, 4, 8}
	quickSort(arr)
	fmt.Println(arr)
}
