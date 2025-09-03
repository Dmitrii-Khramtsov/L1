package main

import "fmt"

func swap(a, b int) (int, int) {
	a = a ^ b
	b = b ^ a
	a = a ^ b
	return a, b
}

func main() {
	a, b := 1, 2
	fmt.Println(swap(a, b))
}
