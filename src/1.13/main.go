package main

import "fmt"

func swap(a, b int) (int, int) {
	fmt.Printf("a = %0.4b\n",a)
	fmt.Printf("b = %0.4b\n\n",b)
	a = a ^ b
	fmt.Printf("a = %0.4b\n",a)
	fmt.Printf("b = %0.4b\n\n",b)
	b = b ^ a
	fmt.Printf("a = %0.4b\n",a)
	fmt.Printf("b = %0.4b\n\n",b)
	a = a ^ b
	fmt.Printf("a = %0.4b\n",a)
	fmt.Printf("b = %0.4b\n\n",b)
	return a, b
}

func main() {
	a, b := 8, 9
	fmt.Println(swap(a, b))
}
