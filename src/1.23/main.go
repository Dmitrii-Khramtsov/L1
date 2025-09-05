package main

import "fmt"

func deletePtrElem(arr []int, i int) []int {
	copy(arr[i:], arr[i+1:])
	arr[len(arr)-1] = 0
	arr = arr[:len(arr)-1]
	slace := make([]int, len(arr)-1)
	slace = arr
	return slace
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	i := 3
	arr2 := deletePtrElem(arr, i)
	fmt.Println(arr2)
	fmt.Println(len(arr2), cap(arr2))
}
