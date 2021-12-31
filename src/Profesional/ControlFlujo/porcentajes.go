package main

import "fmt"

func main() {
	var consumo, descuento float64
	var datosDescuentos string

	fmt.Print("Ingrese consumo: ")
	fmt.Scanln(&consumo)

	if (consumo >= 0) {
		if (consumo <= 100) {
			datosDescuentos = "10%"
			descuento = consumo * 0.10
		} else if (consumo > 100 && consumo <= 200) {
			datosDescuentos = "20%"
			descuento = consumo * 0.20
		} else if (consumo > 200) {
			datosDescuentos = "30%"
			descuento = consumo * 0.30
		}

		montoDescuento := consumo - descuento
		igv := montoDescuento * 0.19
		totalPago := montoDescuento + igv

		fmt.Println("*************************** ")
		fmt.Println("Desto aplicado: ", datosDescuentos)
		fmt.Println("Valor descuento: ", descuento)
		fmt.Println("Valor consumo: ", consumo )
		fmt.Println("Pago IGV: ", igv)
		fmt.Println("Total a pagar: ", totalPago)
		fmt.Println("*************************** ")
	} else {
		fmt.Println("Error al ingresar consumo.")
	}
}
