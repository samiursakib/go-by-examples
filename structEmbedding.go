package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num %v", b.num)
}

type container struct {
	base
	str string
}

type describer interface {
	describe() string
}

func structEmbedding() {
	co := container{
		base: base{num: 1},
		str:  "some name",
	}
	fmt.Printf("co={num: %v, str: %v}", co.num, co.str)
	fmt.Println("\nalso num:", co.base.num)
	fmt.Println("describe:", co.base.describe(), co.describe())

	var d describer = co
	fmt.Println(d.describe())
}
