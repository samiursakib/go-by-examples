package main

import "fmt"

func sums(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func variadicFunctions() {
	args := []int{10, 20, 30}
	sums(1, 2, 3)
	sums(1, 2, 3, 4, 5)
	sums(args...)
}
