package main

import (
	"fmt"
)

func main() {

	for i:=0; i < 10; i++ {
		func (i int) {
			fmt.Println("Into Function", i)
		}(i)
	}

	var number int64;
	number = 10;

//	var p * int64

	p := func (n * int64) * int64 {
		var i int64
		for i = 0; i < 10; i++ {
			* n += i
		}
		return n
	} (&number)

	fmt.Println(number)
	fmt.Println(* p)
}
