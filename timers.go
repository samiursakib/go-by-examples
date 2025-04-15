package main

import (
	"fmt"
	"time"
)

func timers() {
	time1 := time.NewTimer(2 * time.Second)

	<- time1.C
	fmt.Println("timer 1 fired")

	time2 := time.NewTimer(time.Second)
	go func() {
		<- time2.C
		fmt.Println("timer 2 fired")
	}()
	stop2 := time2.Stop()
	if stop2 {
		fmt.Println("timer 2 stopped")
	}

	time.Sleep(2 * time.Second)
}
