package main

import "fmt"

func rangeOverChannels() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for ch := range queue {
		fmt.Println(ch)
	}
}
