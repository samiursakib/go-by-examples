package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func multipleReturnValues() {
	first, second := vals()
	fmt.Println(first, second)
}
