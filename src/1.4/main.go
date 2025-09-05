package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"os"
	"os/signal"
	"sync"
	"time"
)

const nJobs = 10

func worker(ctx context.Context, id int, jobs chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d остановлен: %v\n", id, ctx.Err())
			return
		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("Worker %d завершил: канал закрыт\n", id)
				return
			}
			fmt.Printf("Worker %d обрабатывает job %d\n", id, job)
			time.Sleep(50 * time.Millisecond) // псевдоработа
		}
	}
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	jobs := make(chan int)
	wg := sync.WaitGroup{}

	for i := range nJobs {
		wg.Add(1)
		go worker(ctx, i, jobs, &wg)
	}

	go func() {
		defer close(jobs)
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Генерация заданий остановлена")
				return
			case jobs <- rand.IntN(100):
				time.Sleep(20 * time.Millisecond)
			}
		}
	}()

	wg.Wait()
	fmt.Println("Завершено по Ctrl+c")
}
