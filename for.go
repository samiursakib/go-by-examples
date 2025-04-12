package main

import "fmt"

func for_() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	for j := 5; j < 8; j++ {
		fmt.Println(j)
	}

	for i := range 3 {
		fmt.Println(i)
	}

	for {
		fmt.Println(i)
		i++
		if i == 10 {
			break
		}
	}

	for {
		i++
		if i%2 == 1 {
			continue
		}
		fmt.Println(i)
		if i > 20 {
			break
		}
	}
}
