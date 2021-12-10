package main

import "fmt"

func main() {

	func() {
		fmt.Println("vim-go\n")
	}()

	for i := 0; i <= 10; i++ {
		func(i int) {
			fmt.Println("Hola")
		}(i)
	}

	myFunc := func(nombre string) string {
		return fmt.Sprintf("\nHola, %s", nombre)
	}

	fmt.Println(myFunc)
	fmt.Println(myFunc("Giovanny"))
}
