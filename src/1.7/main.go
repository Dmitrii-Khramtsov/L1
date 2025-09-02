// 1.7
package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	mu sync.Mutex
	rw sync.RWMutex
	m  map[string]int
}

func newSafeMap() *SafeMap {
	return &SafeMap{m: make(map[string]int, 100)}
}

func (s *SafeMap) Set(key string, value int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[key] = value
}

func (s *SafeMap) Get(key string) (int, bool) {
	s.rw.RLock()
	defer s.rw.RUnlock()
	res, ok := s.m[key]
	return res, ok
}

func main() {
	sm := newSafeMap()
	var wg sync.WaitGroup

	for i := range 100 {
		wg.Add(1)
		go func() {
			sm.Set(fmt.Sprintf("key_%d", i), i)
			wg.Done()
		}()
	}

	wg.Wait()

	for i := range 10 {
		if val, ok := sm.Get(fmt.Sprintf("key_%d", i)); ok {
			fmt.Println(val)
		}
	}
}
