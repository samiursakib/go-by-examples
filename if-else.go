package main

import "fmt"

func ifElse() {
	if 7%2 == 1 {
		fmt.Println("7 is odd")
	} else {
		fmt.Println("7 is even")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("Either 8 or 7 is even")
	}

	if num := 9; num < 0 {
		fmt.Println("Number is negative")
	} else if num < 10 {
		fmt.Println("Number has 1 digit")
	} else {
		fmt.Println("Number has multiple digits")
	}
}
