package main

import (
	"errors"
	"fmt"
)

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg, nil
}

var ErrOutOfTea = fmt.Errorf("no more tea available")
var ErrPower = fmt.Errorf("can't boil water")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {
		return fmt.Errorf("making tea: %w", ErrPower)
	}
	return nil
}

func errors_() {
	for _, i := range []int{7, 42} {
		if r, e := f(i); e != nil {
			fmt.Println("error:", e)
		} else {
			fmt.Println("worked:", r)
		}
	}

	for i := range 5 {
		if err := makeTea(i); err != nil {
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("We should buy new Tea")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("Now it is dark")
			} else {
				fmt.Printf("Unknown error: %s", err)
			}
		}
	}
}
