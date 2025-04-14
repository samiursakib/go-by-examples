package main

import (
	"fmt"
	"iter"
	"slices"
)

type List_[T any] struct {
	head, tail *element_[T]
}

type element_[T any] struct {
	next *element_[T]
	val  T
}

func (lst *List_[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element_[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element_[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List_[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1
		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func rangeOverIterators() {
	lst := List_[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	for e := range lst.All() {
		fmt.Println(e)
	}

	all := slices.Collect(lst.All())
	fmt.Println("all:", all)

	for n := range genFib() {
		if n >= 10 {
			break
		}
		fmt.Println(n)
	}
}
