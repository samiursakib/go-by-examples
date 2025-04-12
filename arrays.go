package main

import "fmt"

func arrays() {
	var a [5]int
	a[4] = 100
	fmt.Println(a, "length:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	c := [...]int{10, 11, 12}
	fmt.Println(c, "length:", len(c))

	d := [...]int{100, 200, 30: 400, 500}
	fmt.Println(d)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD[1], twoD)

	anotherTwoD := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(anotherTwoD)
}
