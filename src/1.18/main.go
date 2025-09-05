package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu      sync.Mutex
	counter int32
}

// atomic.AddInt32(&c.counter, 1)
func (c *Counter) Inc() {
	c.mu.Lock()
	c.counter++
	c.mu.Unlock()
}

func (c *Counter) Value() int32 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.counter
}

func (c *Counter) concurrencyCounter() {
	var wg sync.WaitGroup
	for range 1000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}

	wg.Wait()
}

func main() {
	c := &Counter{}
	c.concurrencyCounter()
	fmt.Println(c.Value())
}
