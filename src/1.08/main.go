package main

import "fmt"

func setBit(num int64, iBit, bit int) int64 {
	mask := int64(1) << (iBit - 1)

	switch bit {
	case 1:
		num |= mask
	case 0:
		num &^= mask
	default:
		fmt.Println("not bit num")
	}

	return num
}

func main() {
	var num int64 = 5
	var iBit, bit int = 1, 0

	res := setBit(num, iBit, bit)

	fmt.Printf("num = %032b\n", num)
	fmt.Printf("res = %032b\n", res)
}
