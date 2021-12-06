package main

import "fmt"

func calcular(v float32) float32 {
	return (v * 18) / 100
}

func main() {
	var igv, venta float32

	fmt.Print("Ingresar el precio de venta: ")
	fmt.Scan(&venta)

	igv = calcular(venta)

	fmt.Println("EL IGV de venta es:", igv)
}
