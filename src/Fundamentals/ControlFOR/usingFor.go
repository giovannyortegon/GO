package main

import "fmt"

func main() {
	hello:= "hola Golang"

	for i:= 0; i <= 10; i++ {
		fmt.Println(i,string(hello[i]))
	}
}
