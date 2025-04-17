package main

import (
	"fmt"
	"sync"
	"time"
)

func workerForWaitGroups(id int) {
	fmt.Printf("worker %d starting\n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("worker %d done\n", id)
}

func waitGroups() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			workerForWaitGroups(i)
		}()
	}
	wg.Wait()
}
