package main

import "fmt"

func flippingStr(str string) string {
	reversStr := make([]rune, len(str))
	j := 0
	for i := len([]rune(str))-1; i >= 0; i-- {
		reversStr[j] = []rune(str)[i]
		j++
	}
	return string(reversStr)
}

func main() {
	str := "okey 💡 а роза упала на лапу азора"
	fmt.Println(flippingStr(str))
}
