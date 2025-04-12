package main

import "fmt"

func rangeOverBuiltinTypes() {
	nums := []int{2, 3, 4}
	var sum int
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	for index, num := range nums {
		if num == 3 {
			fmt.Println("index is", index)
		}
	}

	fruits := map[string]string{"a": "apple", "b": "banana"}
	for key, fruit := range fruits {
		fmt.Println("key:", key, "fruit:", fruit)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
