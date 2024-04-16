package main

import (
	"fmt"
	"math"
)

func main() {
	var u int     // u==contaCliente
	var x float64 // x==metrosCubicosGastos
	// v==valorConsumo

	var tipoConta string
	// tipoConsumidor: R: Residencial
	// tipoConsumidor: C: Comercial
	// tipoConsumidor: I: Industrial

	//USAR SWITCH

	fmt.Scanln(&u, &x, &tipoConta)

	if tipoConta == "R" {
		var vR float64 = (5 + 0.05*x)
		fmt.Println("CONTA = ", u)
		fmt.Println("VALOR DA CONTA = ", vR)
	}

	if tipoConta == "C" {
		if x <= 80 {
			fmt.Println("CONTA = ", u)
			fmt.Println("O VALOR DA CONTA = 500.00")
		} else {
			var vC float64 = (500 + 0.25*(x-80))
			fmt.Println("CONTA = ", u)
			fmt.Println("O VALOR DA CONTA = ", math.Round(vC*100)/(100))
		}
	}

	if tipoConta == "I" {
		if x <= 100 {
			fmt.Println("CONTA = ", u)
			fmt.Println("O VALOR DA CONTA = 800.00")
		} else {
			var vI float64 = (800 + 0.04*(x-100))
			fmt.Println("CONTA = ", u)
			fmt.Println("O VALOR DA CONTA = ", math.Round(vI*100)/(100))
		}
	}
}
