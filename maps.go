package main

import (
	"fmt"
	"maps"
)

func maps_() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 8
	m["k3"] = 9
	fmt.Println(m, "value 2:", m["k2"])
	fmt.Println("if key doesn't exist, so default value:", m["v2"])
	fmt.Println("length:", len(m))

	delete(m, "k2")
	fmt.Println(m)
	clear(m)
	fmt.Println(m)

	_, exists := m["v2"]
	fmt.Println(exists)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n and n2 both are equal")
	} else {
		fmt.Println("n and n2 both are not equal")
	}
}
