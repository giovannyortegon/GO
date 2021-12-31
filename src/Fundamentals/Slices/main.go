package main

import (
	"fmt"
	"sort"
)

func main() {

	s1 := make([]int, 5)
	fmt.Println(s1)
	s1 = append(s1, -1)
	fmt.Println(s1)
	s1 = append(s1, -2)
	fmt.Println(s1)
	s1 = append(s1, -3)
	fmt.Println(s1)
	reSlice := s1[1:3]
	fmt.Println("Reslice: ")
	fmt.Println(reSlice)

	a1 := []int{1,2,3,4,5}
	a2 := []int{6,7,8,9,0}

	fmt.Println(a1)
	fmt.Println(a2)
	copy(a1, a2[3:])
	fmt.Println("Copy()", sort.a1)

	var s = make([]string, 0)
	s = append(s, "some string")
	fmt.Println(s)
}
