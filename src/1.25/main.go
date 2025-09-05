package main

import (
	"fmt"
	"time"
)

func mySleep(d time.Duration) {
	<-time.After(d)
}

func main() {
	fmt.Println("Start")
	mySleep(2 * time.Second)
	fmt.Println("End")
}
