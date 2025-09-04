package main

import "fmt"

func checkType(whatType interface{}) {
	switch v := whatType.(type) {
	case int:
		fmt.Println("int:", v)
	case string:
		fmt.Println("string:", v)
	case bool:
		fmt.Println("bool:", v)
	case chan int:
		fmt.Println("chan int:", v)
	case chan string:
		fmt.Println("chan string:", v)
	case chan bool:
		fmt.Println("chan bool:", v)
	default:
		fmt.Println("unknown type")
	}
}

func main() {
	ch := make(chan int, 1)
	ch <- 1
	close(ch)
	ch2 := make(chan string, 1)
	ch2 <- "1"
	close(ch2)
	ch3 := make(chan bool, 1)
	ch3 <- true
	close(ch3)
	checkType(42)
	checkType("good")
	checkType(true)
	checkType(ch)
	checkType(ch2)
	checkType(ch3)
}
