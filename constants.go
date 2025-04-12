package main

import (
	"fmt"
	"math"
)

func constants() {
	const s string = "constants"
	fmt.Println(s)
	const n = 5e7
	fmt.Println(n)
	const res = 3e20 / n
	fmt.Println(res)
	fmt.Println(int64(res), int64(n))
	var signedValue = math.Sin(n)
	fmt.Println(signedValue)
}
