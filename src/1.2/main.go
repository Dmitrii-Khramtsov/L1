package main

import (
	"fmt"
	"sync"
)

func funOut(in []int) <-chan int {
	ch := make(chan int)
	wg := sync.WaitGroup{}

	for _, n := range in {
		wg.Add(1)
		go func(num int) {
			ch <- num * num
			wg.Done()
		}(n)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}

func main() {
	arr := []int{2, 4, 6, 8, 10}

	ch := funOut(arr)

	for n := range ch {
		fmt.Println(n)
	}
}
