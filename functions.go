package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusplus(a, b, c int) int {
	return a + b + c
}

func functions() {
	res := plus(2, 3)
	fmt.Println(res)

	res2 := plusplus(10, 20, 30)
	fmt.Println(res2)
}
