package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func timer(stopped *atomic.Bool) {
	time.Sleep(3000 * time.Millisecond)
	stopped.Store(true)
}

func worckerWithFlag(wg *sync.WaitGroup, stopped *atomic.Bool) {
	defer wg.Done()
	for {
		if stopped.Load() {
			fmt.Println("Воркер worckerWithFlag останавливается")
			return
		}
		fmt.Println("Работает worckerWithFlag...")
		time.Sleep(1000 * time.Millisecond)
	}
}

func timerChan(stoppedChan chan struct{}) {
	time.Sleep(3000 * time.Millisecond)
	stoppedChan <- struct{}{}
}

func worckerSignalChan(wg *sync.WaitGroup, stoppedChan chan struct{}) {
	defer wg.Done()
	for {
		select {
		case <-stoppedChan:
			fmt.Println("Воркер worckerSignalChan останавливается")
			return
		default:
			fmt.Println("Работает worckerSignalChan...")
			time.Sleep(1000 * time.Millisecond)
		}
	}
}

func worckerSignalWithCloseChan(wg *sync.WaitGroup, stoppedChan chan struct{}) {
	defer wg.Done()
	ticker := time.NewTicker(1000 * time.Millisecond)
	for {
		select {
		case _, ok := <-stoppedChan:
			if !ok {
				fmt.Println("Воркер worckerSignalWithCloseChan останавливается")
				return
			}
		case <-ticker.C:
			fmt.Println("Работает worckerSignalWithCloseChan...")
			time.Sleep(1000 * time.Millisecond)
		}
	}
}

func worckerWithContext(wg *sync.WaitGroup, stoppedContext context.Context) {
	defer wg.Done()
	ticker := time.NewTicker(1000 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			fmt.Println("Работает worckerWithContext...")
			time.Sleep(1000 * time.Millisecond)
		case <-stoppedContext.Done():
			fmt.Println("Воркер worckerWithContext останавливается")
			return
		}
	}
}

func workerWithTimeout(wg *sync.WaitGroup) {
	defer wg.Done()
	timeout := time.After(3000 * time.Millisecond)
	for {
		select {
		case <-timeout:
			fmt.Println("Воркер workerWithTimeout останавливается")
			return
		default:
			fmt.Println("Работает workerWithTimeout...")
			time.Sleep(1000 * time.Millisecond)
		}
	}
}

func main() {
	wg := sync.WaitGroup{}

	stopped := atomic.Bool{}
	wg.Add(1)
	go worckerWithFlag(&wg, &stopped)
	go timer(&stopped)

	stoppedChan := make(chan struct{}, 1)
	wg.Add(1)
	go worckerSignalChan(&wg, stoppedChan)
	go timerChan(stoppedChan)

	stoppedChanClosed := make(chan struct{}, 1)
	wg.Add(1)
	go worckerSignalWithCloseChan(&wg, stoppedChanClosed)
	go func() {
		time.Sleep(3000 * time.Millisecond)
		close(stoppedChanClosed)
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 3000*time.Millisecond)
	defer cancel()

	wg.Add(1)
	go worckerWithContext(&wg, ctx)

	wg.Add(1)
	go workerWithTimeout(&wg)

	wg.Wait()
}
