package main

import (
	"fmt"
	"math/rand/v2"
)

func numGenerator() chan int {
	ch := make(chan int, 100)
	go func() {
		for range 100 {
			ch <- rand.IntN(100)
		}
		close(ch)
	}()

	return ch
}

func numModification(in <-chan int) <-chan int {
	out := make(chan int, 100)
	go func() {
		for i := range in {
			out <- i * 2
		}
		close(out)
	}()
	return out
}

func main() {
	ch1 := numGenerator()
	ch2 := numModification(ch1)

	for num := range ch2 {
		fmt.Println(num)
	}
}
