package main

import (
	"fmt"
	"time"
)

func f() {
	fmt.Println("f fucntion")
}

func main() {
	go f()
	time.Sleep(1 * time.Second)
	fmt.Println("main function")
}
