package main

import (
	"fmt"
	"time"
)

func mySleep1(d time.Duration) {
	done := make(chan struct{})

	go func() {
		<-time.After(d)
		close(done)
	}()

	<-done
}

func mySleep2(t time.Duration) {
	done := make(chan struct{})
	go func() {
		start := time.Now()
		for time.Since(start) < t {
		}
		close(done)
	}()
	<-done
}

func mySleep3(t time.Duration) {
	select {
	case <-time.After(t):
	}
}

func main() {
	fmt.Println("Start")
	mySleep1(2 * time.Second)
	fmt.Println("End1")
	mySleep2(2 * time.Second)
	fmt.Println("End2")
	mySleep3(2 * time.Second)
	fmt.Println("End3")
}
