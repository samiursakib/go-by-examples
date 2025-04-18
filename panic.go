package main

import "os"

func panic_() {
	panic("a problem")

	_, err := os.Create("./temp.txt")
	if err != nil {
		panic(err)
	}
}
