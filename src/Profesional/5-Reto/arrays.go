package main

import (
	"fmt"
)

func main() {
	arr := []int{ 1,2,3}

	for i := range arr {
		fmt.Println(i)
	}
}
