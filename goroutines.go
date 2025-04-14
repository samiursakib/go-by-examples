package main

import (
	"fmt"
	"time"
)

func h(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func goroutines() {

	h("direct")

	go h("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
