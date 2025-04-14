package main

import "fmt"

func slices_() {
	var s []string
	fmt.Println(s, "length:", len(s), "capacity:", cap(s))

	s = make([]string, 3)
	fmt.Println(s, "length:", len(s), "capacity:", cap(s))

	s[0] = "a"
	s[1] = "b"
	fmt.Println("value:", s[2], "empty:", s[2] == "")
	s[2] = "c"
	fmt.Println(s, "length:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println(s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copied slice c:", c)

	fmt.Println("slice of slices:", "upto 4", c[:5], "from 2", c[1:], "from 2 upto 4", c[2:5])

	d := []string{"g", "h", "i", "j", "k"}
	fmt.Println(d)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLength := i + 1
		twoD[i] = make([]int, innerLength)
		for j := 0; j < innerLength; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}
