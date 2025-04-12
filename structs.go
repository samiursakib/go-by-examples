package main

import "fmt"

type Person struct {
	name string
	age  int
}

func newPerson(name string) *Person {
	p := Person{name: name}
	p.age = 42
	return &p
}

func structs() {
	fmt.Println(Person{"Bob", 20})
	fmt.Println(Person{name: "Alice", age: 30})
	fmt.Println(&Person{name: "Alice", age: 30})
	fmt.Println(Person{name: "Dogus"})
	fmt.Println(Person{age: 50})

	s := Person{name: "Tom"}
	s.age = 20
	fmt.Println(s, s.name, s.age)

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}
