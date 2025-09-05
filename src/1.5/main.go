package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()

	jobs := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func(j chan int) {
		defer wg.Done()
		id := 0
		for {
			select {
			case job, ok := <-j:
				if !ok {
					fmt.Println("Канал закрыт, воркер завершён")
					return
				}
				fmt.Printf("Worker %d обрабатывает job %d\n", id, job)
				time.Sleep(20 * time.Millisecond) // псевдоработа
			case <-ctx.Done():
				fmt.Printf("Worker %d остановлен: %v\n", id, ctx.Err())
				return
			}
			id++
		}
	}(jobs)

	go func() {
		defer close(jobs)
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Таймер сработал")
				return
			case jobs <- rand.Intn(100):
				time.Sleep(10 * time.Millisecond)
			}
		}
	}()

	wg.Wait()
}
