package main

import (
	"fmt"
)

func goslices() {
	s := make([]string, 3)
	fmt.Println("emp: ", s)
	fmt.Println("len: ", len(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("emp: ", s)
	fmt.Println("len: ", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")

	fmt.Println("emp: ", s)
	fmt.Println("len: ", len(s))

	c := make([]string, len(s))
	copy(c, s)

	fmt.Println(c)

	l := s[2:5]
	fmt.Println("l: ", l)
	s[3] = "xxx"
	fmt.Println("l: ", l)

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
